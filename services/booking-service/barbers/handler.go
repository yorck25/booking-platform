package barbers

import (
	"booking-service/core"
	"strings"
	"time"

	"github.com/google/uuid"
)

func HandleGetReservationSlots(ctx *core.WebContext) error {
	barberIDStr := strings.TrimSpace(ctx.QueryParam("barberId"))
	serviceIDStr := strings.TrimSpace(ctx.QueryParam("serviceId"))
	bookingDateStr := strings.TrimSpace(ctx.QueryParam("bookingDate"))

	if barberIDStr == "" || serviceIDStr == "" || bookingDateStr == "" {
		return ctx.BadRequest("barberId, serviceId and bookingDate are required")
	}

	barberID, err := uuid.Parse(barberIDStr)
	if err != nil {
		return ctx.BadRequest("barberId is not a valid uuid")
	}

	serviceID, err := uuid.Parse(serviceIDStr)
	if err != nil {
		return ctx.BadRequest("serviceId is not a valid uuid")
	}

	bookingDate, err := time.Parse("2006-01-02", bookingDateStr)
	if err != nil {
		return ctx.BadRequest("bookingDate must be in format YYYY-MM-DD")
	}

	req := GetReservationSlotsRequest{
		BarberID:    barberID,
		ServiceID:   serviceID,
		BookingDate: bookingDate,
	}

	store := NewStore(ctx)

	slots, err := store.GetReservationSlots(req)
	if err != nil {
		return ctx.InternalError("fail to get reservation slots: " + err.Error())
	}

	return ctx.Success(slots)
}
