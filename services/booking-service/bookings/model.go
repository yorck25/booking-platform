package bookings

import (
	"time"

	"github.com/google/uuid"
)

type CreateBookingRequest struct {
	BarberID uuid.UUID `json:"barberId" db:"barber_id"`

	CustomerFirstName   string  `json:"customerFirstName" db:"customer_first_name"`
	CustomerLastName    string  `json:"customerLastName" db:"customer_last_name"`
	CustomerPhoneNumber string  `json:"customerPhoneNumber" db:"customer_phone_number"`
	CustomerEmail       *string `json:"customerEmail,omitempty" db:"customer_email"`

	BookingDate time.Time `json:"bookingDate" db:"booking_date"`
	StartTime   time.Time `json:"startTime" db:"start_time"`

	Notes *string `json:"notes,omitempty" db:"notes"`
}

type Booking struct {
	ID                  uuid.UUID  `json:"id" db:"id"`
	BarberID            uuid.UUID  `json:"barberId" db:"barber_id"`
	CustomerFirstName   string     `json:"customerFirstName" db:"customer_first_name"`
	CustomerLastName    string     `json:"customerLastName" db:"customer_last_name"`
	CustomerPhoneNumber string     `json:"customerPhoneNumber" db:"customer_phone_number"`
	CustomerEmail       *string    `json:"customerEmail,omitempty" db:"customer_email"`
	BookingDate         time.Time  `json:"bookingDate" db:"booking_date"`
	StartTime           time.Time  `json:"startTime" db:"start_time"`
	Status              string     `json:"status" db:"status"`
	CancelReason        *string    `json:"cancelReason,omitempty" db:"cancel_reason"`
	Notes               *string    `json:"notes,omitempty" db:"notes"`
	ConfirmedAt         *time.Time `json:"confirmedAt,omitempty" db:"confirmed_at"`
	CanceledAt          *time.Time `json:"canceledAt,omitempty" db:"canceled_at"`
	CreatedAt           time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt           time.Time  `json:"updatedAt" db:"updated_at"`
}
