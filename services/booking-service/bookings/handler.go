package bookings

import (
	"booking-service/core"
)

func HandleCreateBooking(ctx *core.WebContext) error {
	var cbr CreateBookingRequest

	if err := ctx.Bind(&cbr); err != nil {
		return ctx.BadRequest("invalid input")
	}

	store := NewStore(ctx)

	err := store.CreateBooking(cbr)
	if err != nil {
		return ctx.InternalError("Fail to create booking: " + err.Error())
	}

	return ctx.Success("CreateBookingSuccess")
}

func CancelBooking(ctx *core.WebContext) error {
	return ctx.NotFound("Not Implemented yet")
}
