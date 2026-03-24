package barbers

import (
	"time"

	"github.com/google/uuid"
)

type Barber struct {
	ID           uuid.UUID `json:"id" db:"id"`
	InternalName string    `json:"internalName" db:"internal_name"`
	DisplayName  string    `json:"displayName" db:"display_name"`
	Email        *string   `json:"email,omitempty" db:"email"`
	Phone        *string   `json:"phone,omitempty" db:"phone"`
	MobilePhone  *string   `json:"mobilePhone,omitempty" db:"mobile_phone"`
	PostalCode   *string   `json:"postalCode,omitempty" db:"postal_code"`
	Street       *string   `json:"street,omitempty" db:"street"`
	City         *string   `json:"city,omitempty" db:"city"`
	Active       bool      `json:"active" db:"active"`
	Deleted      bool      `json:"deleted" db:"deleted"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

type BarberOpeningHours struct {
	ID        uuid.UUID `json:"id" db:"id"`
	BarberID  uuid.UUID `json:"barberId" db:"barber_id"`
	Weekday   int16     `json:"weekday" db:"weekday"`
	StartTime time.Time `json:"startTime" db:"start_time"`
	EndTime   time.Time `json:"endTime" db:"end_time"`
	IsClosed  bool      `json:"isClosed" db:"is_closed"`
}

type BarberBreak struct {
	ID        uuid.UUID `json:"id" db:"id"`
	BarberID  uuid.UUID `json:"barberId" db:"barber_id"`
	Weekday   int16     `json:"weekday" db:"weekday"`
	StartTime time.Time `json:"startTime" db:"start_time"`
	EndTime   time.Time `json:"endTime" db:"end_time"`
}

type ReservationSlot struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	IsBooked  bool   `json:"isBooked"`
}

type GetReservationSlotsRequest struct {
	BarberID    uuid.UUID `json:"barberId"`
	BookingDate time.Time `json:"bookingDate"`
}
