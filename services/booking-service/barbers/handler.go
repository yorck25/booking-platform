package barbers

import (
	"booking-service/core"
	"time"

	"github.com/google/uuid"
)

func HandleGetReservationSlots(ctx *core.WebContext) error {
	barberIDStr := ctx.QueryParam("barberId")
	bookingDateStr := ctx.QueryParam("bookingDate")

	if barberIDStr == "" || bookingDateStr == "" {
		return ctx.BadRequest("barberId and bookingDate are required")
	}

	barberID, err := uuid.Parse(barberIDStr)
	if err != nil {
		return ctx.BadRequest("barberId is not a valid uuid")
	}

	BookingDate, err := time.Parse("2006-01-02", bookingDateStr)
	if err != nil {
		return ctx.BadRequest("bookingDate must be in format YYYY-MM-DD")
	}

	req := GetReservationSlotsRequest{
		BarberID:    barberID,
		BookingDate: BookingDate,
	}

	store := NewStore(ctx)

	slots, err := store.GetReservationSlots(req)
	if err != nil {
		return ctx.InternalError("Fail to get reservation slots: " + err.Error())
	}

	return ctx.Success(slots)
}
