package bookings

import (
	"booking-service/core"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

var (
	ErrServiceNotFound         = errors.New("service not found")
	ErrNoAvailableEmployee     = errors.New("no available employee")
	ErrInvalidDate             = errors.New("invalid date")
	ErrInvalidTime             = errors.New("invalid time")
	ErrBookingNotFound         = errors.New("booking not found")
	ErrBookingCannotBeCanceled = errors.New("booking cannot be canceled")
)

type Store struct {
	db *sqlx.DB
}

func NewStore(ctx *core.WebContext) *Store {
	return &Store{db: ctx.GetDb()}
}

func (s *Store) CreateBooking(req CreateBookingRequest) (*Booking, error) {
	bookingDate, err := time.Parse("2006-01-02", req.BookingDate)
	if err != nil {
		return nil, ErrInvalidDate
	}

	startClock, err := parseClock(req.StartTime)
	if err != nil {
		return nil, ErrInvalidTime
	}

	tx, err := s.db.Beginx()
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	service, err := s.getService(tx, req.BarberID, req.ServiceID)
	if err != nil {
		return nil, err
	}
	if service == nil || !service.Active || service.Deleted {
		return nil, ErrServiceNotFound
	}

	endClock := startClock.Add(time.Duration(service.DurationMinutes) * time.Minute)

	isClosed, err := s.isBarberClosedDay(tx, req.BarberID, bookingDate)
	if err != nil {
		return nil, err
	}
	if isClosed {
		return nil, ErrNoAvailableEmployee
	}

	employees, err := s.getEmployeesForService(tx, req.BarberID, req.ServiceID)
	if err != nil {
		return nil, err
	}
	if len(employees) == 0 {
		return nil, ErrNoAvailableEmployee
	}

	weekday := int16(bookingDate.Weekday())

	var selectedEmployeeID uuid.UUID
	found := false

	for _, emp := range employees {
		ok, err := s.employeeCanTakeSlot(tx, emp.ID, weekday, bookingDate, startClock, endClock)
		if err != nil {
			return nil, err
		}
		if ok {
			selectedEmployeeID = emp.ID
			found = true
			break
		}
	}

	if !found {
		return nil, ErrNoAvailableEmployee
	}

	now := time.Now()
	booking := &Booking{
		ID:                  uuid.New(),
		BarberID:            req.BarberID,
		EmployeeID:          selectedEmployeeID,
		ServiceID:           req.ServiceID,
		CustomerFirstName:   req.CustomerFirstName,
		CustomerLastName:    req.CustomerLastName,
		CustomerPhoneNumber: req.CustomerPhoneNumber,
		CustomerEmail:       req.CustomerEmail,
		BookingDate:         bookingDate,
		StartTime:           startClock.Format("15:04"),
		EndTime:             endClock.Format("15:04"),
		Status:              "pending",
		Notes:               normalizeOptionalString(req.Notes),
		ServiceName:         service.DisplayName,
		ServiceDurationMin:  service.DurationMinutes,
		ServicePriceCents:   service.PriceCents,
		TermsAccepted:       true,
		TermsAcceptedAt:     &now,
		SMSSent:             false,
		CreatedAt:           now,
		UpdatedAt:           now,
	}

	stmt, err := tx.PrepareNamed(`
		INSERT INTO bookings (
			id,
			barber_id,
			employee_id,
			service_id,
			customer_first_name,
			customer_last_name,
			customer_phone_number,
			customer_email,
			booking_date,
			start_time,
			end_time,
			status,
			cancel_reason,
			notes,
			service_name,
			service_duration_min,
			service_price_cents,
			terms_accepted,
			terms_accepted_at,
			confirmed_at,
			rejected_at,
			canceled_at,
			sms_sent,
			sms_sent_at,
			sms_error,
			created_at,
			updated_at
		) VALUES (
			:id,
			:barber_id,
			:employee_id,
			:service_id,
			:customer_first_name,
			:customer_last_name,
			:customer_phone_number,
			:customer_email,
			:booking_date,
			:start_time,
			:end_time,
			:status,
			:cancel_reason,
			:notes,
			:service_name,
			:service_duration_min,
			:service_price_cents,
			:terms_accepted,
			:terms_accepted_at,
			:confirmed_at,
			:rejected_at,
			:canceled_at,
			:sms_sent,
			:sms_sent_at,
			:sms_error,
			:created_at,
			:updated_at
		)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(booking)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return booking, nil
}

func (s *Store) CancelBooking(req CancelBookingRequest) error {
	var status string
	var canceledAt *time.Time

	err := s.db.Get(&status, `
		SELECT status
		FROM bookings
		WHERE id = $1
		  AND customer_phone_number = $2
		LIMIT 1
	`, req.BookingID, req.CustomerPhoneNumber)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrBookingNotFound
		}
		return err
	}

	if status != "pending" && status != "confirmed" {
		return ErrBookingCannotBeCanceled
	}

	now := time.Now()
	canceledAt = &now

	_, err = s.db.Exec(`
		UPDATE bookings
		SET
			status = 'cancelled',
			cancel_reason = $3,
			canceled_at = $4,
			updated_at = NOW()
		WHERE id = $1
		  AND customer_phone_number = $2
	`, req.BookingID, req.CustomerPhoneNumber, normalizeOptionalString(req.CancelReason), canceledAt)

	return err
}

func (s *Store) getService(tx *sqlx.Tx, barberID, serviceID uuid.UUID) (*service, error) {
	var svc service

	err := tx.Get(&svc, `
		SELECT
			id,
			barber_id,
			display_name,
			duration_minutes,
			price_cents,
			active,
			deleted
		FROM services
		WHERE id = $1
		  AND barber_id = $2
		  AND deleted = false
		LIMIT 1
	`, serviceID, barberID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &svc, nil
}

func (s *Store) getEmployeesForService(tx *sqlx.Tx, barberID, serviceID uuid.UUID) ([]employee, error) {
	var employees []employee

	err := tx.Select(&employees, `
		SELECT be.id
		FROM barber_employees be
		INNER JOIN employee_services es ON es.employee_id = be.id
		WHERE be.barber_id = $1
		  AND es.service_id = $2
		  AND be.active = true
		  AND be.deleted = false
		ORDER BY be.created_at ASC
	`, barberID, serviceID)

	return employees, err
}

func (s *Store) employeeCanTakeSlot(
	tx *sqlx.Tx,
	employeeID uuid.UUID,
	weekday int16,
	bookingDate time.Time,
	slotStart time.Time,
	slotEnd time.Time,
) (bool, error) {
	workingHours, err := s.getEmployeeWorkingHours(tx, employeeID, weekday)
	if err != nil {
		return false, err
	}
	if len(workingHours) == 0 {
		return false, nil
	}

	insideWorkingTime := false
	for _, wh := range workingHours {
		if wh.IsClosed {
			continue
		}

		whStart, err := parseClock(wh.StartTime)
		if err != nil {
			continue
		}
		whEnd, err := parseClock(wh.EndTime)
		if err != nil {
			continue
		}

		if (slotStart.Equal(whStart) || slotStart.After(whStart)) &&
			(slotEnd.Equal(whEnd) || slotEnd.Before(whEnd)) {
			insideWorkingTime = true
			break
		}
	}
	if !insideWorkingTime {
		return false, nil
	}

	breaks, err := s.getEmployeeBreaks(tx, employeeID, weekday)
	if err != nil {
		return false, err
	}
	for _, br := range breaks {
		breakStart, err := parseClock(br.StartTime)
		if err != nil {
			continue
		}
		breakEnd, err := parseClock(br.EndTime)
		if err != nil {
			continue
		}

		if slotStart.Before(breakEnd) && slotEnd.After(breakStart) {
			return false, nil
		}
	}

	bookedRanges, err := s.getEmployeeBookedRanges(tx, employeeID, bookingDate)
	if err != nil {
		return false, err
	}
	for _, booked := range bookedRanges {
		bookedStart, err := parseClock(booked.StartTime)
		if err != nil {
			continue
		}
		bookedEnd, err := parseClock(booked.EndTime)
		if err != nil {
			continue
		}

		if slotStart.Before(bookedEnd) && slotEnd.After(bookedStart) {
			return false, nil
		}
	}

	return true, nil
}

func (s *Store) getEmployeeWorkingHours(tx *sqlx.Tx, employeeID uuid.UUID, weekday int16) ([]workingHour, error) {
	var items []workingHour

	err := tx.Select(&items, `
		SELECT
			start_time::text AS start_time,
			end_time::text AS end_time,
			is_closed
		FROM employee_working_hours
		WHERE employee_id = $1
		  AND weekday = $2
		ORDER BY start_time
	`, employeeID, weekday)

	return items, err
}

func (s *Store) getEmployeeBreaks(tx *sqlx.Tx, employeeID uuid.UUID, weekday int16) ([]employeeBreak, error) {
	var items []employeeBreak

	err := tx.Select(&items, `
		SELECT
			start_time::text AS start_time,
			end_time::text AS end_time
		FROM employee_breaks
		WHERE employee_id = $1
		  AND weekday = $2
		  AND active = true
		ORDER BY start_time
	`, employeeID, weekday)

	return items, err
}

func (s *Store) getEmployeeBookedRanges(tx *sqlx.Tx, employeeID uuid.UUID, bookingDate time.Time) ([]bookedRange, error) {
	var items []bookedRange

	err := tx.Select(&items, `
		SELECT
			start_time::text AS start_time,
			end_time::text AS end_time
		FROM bookings
		WHERE employee_id = $1
		  AND booking_date = $2
		  AND status IN ('pending', 'confirmed')
		  AND canceled_at IS NULL
		ORDER BY start_time
	`, employeeID, bookingDate.Format("2006-01-02"))

	return items, err
}

func (s *Store) isBarberClosedDay(tx *sqlx.Tx, barberID uuid.UUID, bookingDate time.Time) (bool, error) {
	var exists bool

	err := tx.Get(&exists, `
		SELECT EXISTS (
			SELECT 1
			FROM barber_closed_days
			WHERE barber_id = $1
			  AND closed_date = $2
		)
	`, barberID, bookingDate.Format("2006-01-02"))

	return exists, err
}

func parseClock(value string) (time.Time, error) {
	value = strings.TrimSpace(value)

	layouts := []string{
		"15:04:05",
		"15:04",
	}

	var lastErr error
	for _, layout := range layouts {
		t, err := time.Parse(layout, value)
		if err == nil {
			return t, nil
		}
		lastErr = err
	}

	return time.Time{}, lastErr
}

func normalizeOptionalString(value *string) *string {
	if value == nil {
		return nil
	}

	v := strings.TrimSpace(*value)
	if v == "" {
		return nil
	}

	return &v
}
