package bookings

import (
	"booking-service/core"
	"strings"
)

func HandleCreateBooking(ctx *core.WebContext) error {
	var req CreateBookingRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest("invalid input")
	}

	req.CustomerFirstName = strings.TrimSpace(req.CustomerFirstName)
	req.CustomerLastName = strings.TrimSpace(req.CustomerLastName)
	req.CustomerPhoneNumber = strings.TrimSpace(req.CustomerPhoneNumber)
	req.BookingDate = strings.TrimSpace(req.BookingDate)
	req.StartTime = strings.TrimSpace(req.StartTime)

	if req.BarberID.String() == "00000000-0000-0000-0000-000000000000" {
		return ctx.BadRequest("barberId is required")
	}

	if req.ServiceID.String() == "00000000-0000-0000-0000-000000000000" {
		return ctx.BadRequest("serviceId is required")
	}

	if req.CustomerFirstName == "" {
		return ctx.BadRequest("customerFirstName is required")
	}

	if req.CustomerLastName == "" {
		return ctx.BadRequest("customerLastName is required")
	}

	if req.CustomerPhoneNumber == "" {
		return ctx.BadRequest("customerPhoneNumber is required")
	}

	if req.BookingDate == "" {
		return ctx.BadRequest("bookingDate is required")
	}

	if req.StartTime == "" {
		return ctx.BadRequest("startTime is required")
	}

	if !req.TermsAccepted {
		return ctx.BadRequest("terms must be accepted")
	}

	store := NewStore(ctx)

	booking, err := store.CreateBooking(req)
	if err != nil {
		switch err {
		case ErrServiceNotFound:
			return ctx.NotFound("service not found")
		case ErrNoAvailableEmployee:
			return ctx.BadRequest("selected slot is no longer available")
		case ErrInvalidDate:
			return ctx.BadRequest("bookingDate must be in format YYYY-MM-DD")
		case ErrInvalidTime:
			return ctx.BadRequest("startTime must be in format HH:MM")
		default:
			return ctx.InternalError("fail to create booking: " + err.Error())
		}
	}

	return ctx.Success(booking)
}

func CancelBooking(ctx *core.WebContext) error {
	var req CancelBookingRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest("invalid input")
	}

	req.CustomerPhoneNumber = strings.TrimSpace(req.CustomerPhoneNumber)

	if req.BookingID.String() == "00000000-0000-0000-0000-000000000000" {
		return ctx.BadRequest("bookingId is required")
	}

	if req.CustomerPhoneNumber == "" {
		return ctx.BadRequest("customerPhoneNumber is required")
	}

	store := NewStore(ctx)

	err := store.CancelBooking(req)
	if err != nil {
		switch err {
		case ErrBookingNotFound:
			return ctx.NotFound("booking not found")
		case ErrBookingCannotBeCanceled:
			return ctx.BadRequest("booking cannot be canceled")
		default:
			return ctx.InternalError("fail to cancel booking: " + err.Error())
		}
	}

	return ctx.Success("booking canceled successfully")
}
