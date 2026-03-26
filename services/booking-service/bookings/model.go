package bookings

import (
	"time"

	"github.com/google/uuid"
)

type CreateBookingRequest struct {
	BarberID  uuid.UUID `json:"barberId"`
	ServiceID uuid.UUID `json:"serviceId"`

	CustomerFirstName   string  `json:"customerFirstName"`
	CustomerLastName    string  `json:"customerLastName"`
	CustomerPhoneNumber string  `json:"customerPhoneNumber"`
	CustomerEmail       *string `json:"customerEmail,omitempty"`

	BookingDate string `json:"bookingDate"` // YYYY-MM-DD
	StartTime   string `json:"startTime"`   // HH:MM

	Notes         *string `json:"notes,omitempty"`
	TermsAccepted bool    `json:"termsAccepted"`
}

type CancelBookingRequest struct {
	BookingID           uuid.UUID `json:"bookingId"`
	CustomerPhoneNumber string    `json:"customerPhoneNumber"`
	CancelReason        *string   `json:"cancelReason,omitempty"`
}

type Booking struct {
	ID         uuid.UUID `json:"id" db:"id"`
	BarberID   uuid.UUID `json:"barberId" db:"barber_id"`
	EmployeeID uuid.UUID `json:"employeeId" db:"employee_id"`
	ServiceID  uuid.UUID `json:"serviceId" db:"service_id"`

	CustomerFirstName   string  `json:"customerFirstName" db:"customer_first_name"`
	CustomerLastName    string  `json:"customerLastName" db:"customer_last_name"`
	CustomerPhoneNumber string  `json:"customerPhoneNumber" db:"customer_phone_number"`
	CustomerEmail       *string `json:"customerEmail,omitempty" db:"customer_email"`

	BookingDate time.Time `json:"bookingDate" db:"booking_date"`
	StartTime   string    `json:"startTime" db:"start_time"`
	EndTime     string    `json:"endTime" db:"end_time"`

	Status       string  `json:"status" db:"status"`
	CancelReason *string `json:"cancelReason,omitempty" db:"cancel_reason"`
	Notes        *string `json:"notes,omitempty" db:"notes"`

	ServiceName        string `json:"serviceName" db:"service_name"`
	ServiceDurationMin int    `json:"serviceDurationMin" db:"service_duration_min"`
	ServicePriceCents  int    `json:"servicePriceCents" db:"service_price_cents"`

	TermsAccepted   bool       `json:"termsAccepted" db:"terms_accepted"`
	TermsAcceptedAt *time.Time `json:"termsAcceptedAt,omitempty" db:"terms_accepted_at"`

	ConfirmedAt *time.Time `json:"confirmedAt,omitempty" db:"confirmed_at"`
	RejectedAt  *time.Time `json:"rejectedAt,omitempty" db:"rejected_at"`
	CanceledAt  *time.Time `json:"canceledAt,omitempty" db:"canceled_at"`

	SMSSent   bool       `json:"smsSent" db:"sms_sent"`
	SMSSentAt *time.Time `json:"smsSentAt,omitempty" db:"sms_sent_at"`
	SMSError  *string    `json:"smsError,omitempty" db:"sms_error"`

	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type service struct {
	ID              uuid.UUID `db:"id"`
	BarberID        uuid.UUID `db:"barber_id"`
	DisplayName     string    `db:"display_name"`
	DurationMinutes int       `db:"duration_minutes"`
	PriceCents      int       `db:"price_cents"`
	Active          bool      `db:"active"`
	Deleted         bool      `db:"deleted"`
}

type employee struct {
	ID uuid.UUID `db:"id"`
}

type workingHour struct {
	StartTime string `db:"start_time"`
	EndTime   string `db:"end_time"`
	IsClosed  bool   `db:"is_closed"`
}

type employeeBreak struct {
	StartTime string `db:"start_time"`
	EndTime   string `db:"end_time"`
}

type bookedRange struct {
	StartTime string `db:"start_time"`
	EndTime   string `db:"end_time"`
}
