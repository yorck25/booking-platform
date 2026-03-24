package barbers

import (
	"booking-service/core"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(ctx *core.WebContext) *Store {
	return &Store{db: ctx.GetDb()}
}

func (s *Store) GetReservationSlots(req GetReservationSlotsRequest) ([]ReservationSlot, error) {
	weekday := int16(req.BookingDate.Weekday())

	openingHours, err := s.getOpeningHours(req.BarberID, weekday)
	if err != nil {
		return nil, err
	}

	if openingHours == nil || openingHours.IsClosed {
		return []ReservationSlot{}, nil
	}

	breaks, err := s.getBreaks(req.BarberID, weekday)
	if err != nil {
		return nil, err
	}

	bookedTimes, err := s.getBookedStartTimes(req.BarberID, req.BookingDate)
	if err != nil {
		return nil, err
	}

	slots := BuildReservationSlots(*openingHours, breaks, bookedTimes)

	return slots, nil
}

func (s *Store) getOpeningHours(barberID uuid.UUID, weekday int16) (*BarberOpeningHours, error) {
	var openingHours BarberOpeningHours

	err := s.db.Get(&openingHours, `
		SELECT 
			id,
			barber_id,
			weekday,
			start_time,
			end_time,
			is_closed
		FROM barber_opening_hours
		WHERE barber_id = $1
		  AND weekday = $2
		LIMIT 1
	`, barberID, weekday)

	if err != nil {
		return nil, err
	}

	return &openingHours, nil
}

func (s *Store) getBreaks(barberID uuid.UUID, weekday int16) ([]BarberBreak, error) {
	var breaks []BarberBreak

	err := s.db.Select(&breaks, `
		SELECT
			id,
			barber_id,
			weekday,
			start_time,
			end_time
		FROM barber_breaks
		WHERE barber_id = $1
		  AND weekday = $2
		ORDER BY start_time
	`, barberID, weekday)

	return breaks, err
}

func (s *Store) getBookedStartTimes(barberID uuid.UUID, bookingDate time.Time) (map[string]bool, error) {
	var booked []struct {
		StartTime time.Time `db:"start_time"`
	}

	err := s.db.Select(&booked, `
		SELECT start_time
		FROM bookings
		WHERE barber_id = $1
		  AND booking_date = $2
		  AND status IN ('pending', 'confirmed')
		  AND canceled_at IS NULL
	`, barberID, bookingDate.Format("2006-01-02"))

	if err != nil {
		return nil, err
	}

	result := make(map[string]bool, len(booked))
	for _, b := range booked {
		result[b.StartTime.Format("15:04")] = true
	}

	return result, nil
}
