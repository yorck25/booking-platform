package bookings

import (
	"booking-service/core"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(ctx *core.WebContext) *Store {
	return &Store{db: ctx.GetDb()}
}

func (s *Store) CreateBooking(cbr CreateBookingRequest) error {
	stmt, err := s.db.PrepareNamed(`
		INSERT INTO bookings (
			id,
			barber_id,
			customer_first_name,
			customer_last_name,
			customer_phone_number,
			customer_email,
			booking_date,
			start_time,
			status,
			created_at,
			updated_at
		) VALUES (
			:id,
			:barber_id,
			:customer_first_name,
			:customer_last_name,
			:customer_phone_number,
			:customer_email,
			:booking_date,
			:start_time,
			:status,
			now(),
			now()
		)
	`)

	if err != nil {
		return err
	}

	params := map[string]any{
		"id":                    uuid.NewUUID(),
		"barber_id":             cbr.BarberID,
		"customer_first_name":   cbr.CustomerFirstName,
		"customer_last_name":    cbr.CustomerLastName,
		"customer_phone_number": cbr.CustomerPhoneNumber,
		"customer_email":        cbr.CustomerEmail,
		"booking_date":          cbr.BookingDate,
		"start_time":            cbr.StartTime,
		"status":                "pending",
	}

	err = stmt.Get(stmt, params)

	return err
}
