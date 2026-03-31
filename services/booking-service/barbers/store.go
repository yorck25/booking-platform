package barbers

import (
	"booking-service/core"
	"database/sql"
	"sort"
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

func (s *Store) GetBarberByID(barberId string) (Barber, error) {
	var barber Barber

	err := s.db.Get(&barber, `
		SELECT *
		FROM barbers
		WHERE id = $1
		  AND deleted = false
		LIMIT 1
	`, barberId)

	if err != nil {
		return barber, err
	}

	return barber, nil
}

func (s *Store) GetReservationSlots(req GetReservationSlotsRequest) ([]ReservationSlot, error) {
	isClosed, err := s.isBarberClosedDay(req.BarberID, req.BookingDate)
	if err != nil {
		return nil, err
	}
	if isClosed {
		return []ReservationSlot{}, nil
	}

	service, err := s.getService(req.BarberID, req.ServiceID)
	if err != nil {
		return nil, err
	}
	if service == nil || !service.Active || service.Deleted {
		return []ReservationSlot{}, nil
	}

	employees, err := s.getEmployeesForService(req.BarberID, req.ServiceID)
	if err != nil {
		return nil, err
	}
	if len(employees) == 0 {
		return []ReservationSlot{}, nil
	}

	weekday := int16(req.BookingDate.Weekday())
	slotMap := map[string]ReservationSlot{}

	for _, employee := range employees {
		workingHours, err := s.getEmployeeWorkingHours(employee.ID, weekday)
		if err != nil {
			return nil, err
		}
		if len(workingHours) == 0 {
			continue
		}

		breaks, err := s.getEmployeeBreaks(employee.ID, weekday)
		if err != nil {
			return nil, err
		}

		bookedRanges, err := s.getEmployeeBookedRanges(employee.ID, req.BookingDate)
		if err != nil {
			return nil, err
		}

		employeeSlots := BuildEmployeeReservationSlots(
			workingHours,
			breaks,
			bookedRanges,
			service.DurationMinutes,
		)

		for _, slot := range employeeSlots {
			key := slot.StartTime + "-" + slot.EndTime
			existing, exists := slotMap[key]
			if !exists {
				slotMap[key] = slot
				continue
			}

			if existing.IsBooked && !slot.IsBooked {
				slotMap[key] = slot
			}
		}
	}

	result := make([]ReservationSlot, 0, len(slotMap))
	for _, slot := range slotMap {
		if !slot.IsBooked {
			result = append(result, slot)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].StartTime < result[j].StartTime
	})

	return result, nil
}

func (s *Store) getService(barberID, serviceID uuid.UUID) (*Service, error) {
	var service Service

	err := s.db.Get(&service, `
		SELECT
			id,
			barber_id,
			internal_name,
			display_name,
			description,
			duration_minutes,
			price_cents,
			active,
			deleted,
			sort_order,
			created_at,
			updated_at
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

	return &service, nil
}

func (s *Store) getEmployeesForService(barberID, serviceID uuid.UUID) ([]BarberEmployee, error) {
	var employees []BarberEmployee

	err := s.db.Select(&employees, `
		SELECT
			be.id,
			be.barber_id,
			be.display_name,
			be.active,
			be.deleted
		FROM barber_employees be
		INNER JOIN employee_services es ON es.employee_id = be.id
		WHERE be.barber_id = $1
		  AND es.service_id = $2
		  AND be.active = true
		  AND be.deleted = false
		ORDER BY be.display_name ASC
	`, barberID, serviceID)

	return employees, err
}

func (s *Store) getEmployeeWorkingHours(employeeID uuid.UUID, weekday int16) ([]EmployeeWorkingHour, error) {
	var workingHours []EmployeeWorkingHour

	err := s.db.Select(&workingHours, `
		SELECT
			id,
			employee_id,
			weekday,
			start_time::text AS start_time,
			end_time::text AS end_time,
			is_closed
		FROM employee_working_hours
		WHERE employee_id = $1
		  AND weekday = $2
		  AND is_closed = false
		ORDER BY start_time
	`, employeeID, weekday)

	return workingHours, err
}

func (s *Store) getEmployeeBreaks(employeeID uuid.UUID, weekday int16) ([]EmployeeBreak, error) {
	var breaks []EmployeeBreak

	err := s.db.Select(&breaks, `
		SELECT
			id,
			employee_id,
			weekday,
			start_time::text AS start_time,
			end_time::text AS end_time,
			description,
			active
		FROM employee_breaks
		WHERE employee_id = $1
		  AND weekday = $2
		  AND active = true
		ORDER BY start_time
	`, employeeID, weekday)

	return breaks, err
}

func (s *Store) getEmployeeBookedRanges(employeeID uuid.UUID, bookingDate time.Time) ([]bookedRange, error) {
	var booked []bookedRange

	err := s.db.Select(&booked, `
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

	return booked, err
}

func (s *Store) isBarberClosedDay(barberID uuid.UUID, bookingDate time.Time) (bool, error) {
	var exists bool

	err := s.db.Get(&exists, `
		SELECT EXISTS (
			SELECT 1
			FROM barber_closed_days
			WHERE barber_id = $1
			  AND closed_date = $2
		)
	`, barberID, bookingDate.Format("2006-01-02"))

	return exists, err
}
