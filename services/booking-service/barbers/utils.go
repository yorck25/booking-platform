package barbers

import "time"

const slotStepMinutes = 15

func BuildEmployeeReservationSlots(
	workingHours []EmployeeWorkingHour,
	breaks []EmployeeBreak,
	bookedRanges []bookedRange,
	serviceDurationMinutes int,
) []ReservationSlot {
	slots := make([]ReservationSlot, 0)

	for _, wh := range workingHours {
		windowStart, err := parseClock(wh.StartTime)
		if err != nil {
			continue
		}

		windowEnd, err := parseClock(wh.EndTime)
		if err != nil {
			continue
		}

		current := windowStart
		for {
			slotEnd := current.Add(time.Duration(serviceDurationMinutes) * time.Minute)
			if slotEnd.After(windowEnd) {
				break
			}

			isBooked := false

			if IsInEmployeeBreak(current, slotEnd, breaks) {
				isBooked = true
			}

			if !isBooked && IsOverlappingBookedRange(current, slotEnd, bookedRanges) {
				isBooked = true
			}

			slots = append(slots, ReservationSlot{
				StartTime: current.Format("15:04"),
				EndTime:   slotEnd.Format("15:04"),
				IsBooked:  isBooked,
			})

			current = current.Add(slotStepMinutes * time.Minute)
		}
	}

	return slots
}

func IsInEmployeeBreak(slotStart, slotEnd time.Time, breaks []EmployeeBreak) bool {
	for _, employeeBreak := range breaks {
		breakStart, err := parseClock(employeeBreak.StartTime)
		if err != nil {
			continue
		}

		breakEnd, err := parseClock(employeeBreak.EndTime)
		if err != nil {
			continue
		}

		if slotStart.Before(breakEnd) && slotEnd.After(breakStart) {
			return true
		}
	}

	return false
}

func IsOverlappingBookedRange(slotStart, slotEnd time.Time, bookedRanges []bookedRange) bool {
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
			return true
		}
	}

	return false
}

func parseClock(value string) (time.Time, error) {
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
