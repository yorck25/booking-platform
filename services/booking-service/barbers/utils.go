package barbers

import "time"

func BuildReservationSlots(
	openingHours BarberOpeningHours,
	breaks []BarberBreak,
	bookedTimes map[string]bool,
) []ReservationSlot {
	var slots []ReservationSlot

	current := openingHours.StartTime
	end := openingHours.EndTime

	for current.Before(end) {
		slotEnd := current.Add(time.Hour)

		if slotEnd.After(end) {
			break
		}

		if IsInBreak(current, slotEnd, breaks) {
			current = current.Add(time.Hour)
			continue
		}

		startFormatted := current.Format("15:04")
		endFormatted := slotEnd.Format("15:04")

		slots = append(slots, ReservationSlot{
			StartTime: startFormatted,
			EndTime:   endFormatted,
			IsBooked:  bookedTimes[startFormatted],
		})

		current = current.Add(time.Hour)
	}

	return slots
}

func IsInBreak(slotStart, slotEnd time.Time, breaks []BarberBreak) bool {
	for _, barberBreak := range breaks {
		if slotStart.Before(barberBreak.EndTime) && slotEnd.After(barberBreak.StartTime) {
			return true
		}
	}

	return false
}
