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

type Service struct {
	ID              uuid.UUID `json:"id" db:"id"`
	BarberID        uuid.UUID `json:"barberId" db:"barber_id"`
	InternalName    string    `json:"internalName" db:"internal_name"`
	DisplayName     string    `json:"displayName" db:"display_name"`
	Description     string    `json:"description" db:"description"`
	DurationMinutes int       `json:"durationMinutes" db:"duration_minutes"`
	PriceCents      int       `json:"priceCents" db:"price_cents"`
	Active          bool      `json:"active" db:"active"`
	Deleted         bool      `json:"deleted" db:"deleted"`
	SortOrder       int       `json:"sortOrder" db:"sort_order"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}

type BarberEmployee struct {
	ID          uuid.UUID `json:"id" db:"id"`
	BarberID    uuid.UUID `json:"barberId" db:"barber_id"`
	DisplayName string    `json:"displayName" db:"display_name"`
	Active      bool      `json:"active" db:"active"`
	Deleted     bool      `json:"deleted" db:"deleted"`
}

type EmployeeWorkingHour struct {
	ID         uuid.UUID `json:"id" db:"id"`
	EmployeeID uuid.UUID `json:"employeeId" db:"employee_id"`
	Weekday    int16     `json:"weekday" db:"weekday"`
	StartTime  string    `json:"startTime" db:"start_time"`
	EndTime    string    `json:"endTime" db:"end_time"`
	IsClosed   bool      `json:"isClosed" db:"is_closed"`
}

type EmployeeBreak struct {
	ID          uuid.UUID `json:"id" db:"id"`
	EmployeeID  uuid.UUID `json:"employeeId" db:"employee_id"`
	Weekday     int16     `json:"weekday" db:"weekday"`
	StartTime   string    `json:"startTime" db:"start_time"`
	EndTime     string    `json:"endTime" db:"end_time"`
	Description string    `json:"description" db:"description"`
	Active      bool      `json:"active" db:"active"`
}

type ReservationSlot struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	IsBooked  bool   `json:"isBooked"`
}

type GetReservationSlotsRequest struct {
	BarberID    uuid.UUID `json:"barberId"`
	ServiceID   uuid.UUID `json:"serviceId"`
	BookingDate time.Time `json:"bookingDate"`
}

type bookedRange struct {
	StartTime string `db:"start_time"`
	EndTime   string `db:"end_time"`
}
