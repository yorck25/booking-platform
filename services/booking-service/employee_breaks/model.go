package employee_breaks

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeBreak struct {
	Id          uuid.UUID `json:"id" db:"id"`
	EmployeeId  uuid.UUID `json:"employeeId" db:"employee_id"`
	Weekday     int       `json:"weekday" db:"weekday"`
	StartTime   string    `json:"startTime" db:"start_time"`
	EndTime     string    `json:"endTime" db:"end_time"`
	Description string    `json:"description" db:"description"`
	Active      bool      `json:"active" db:"active"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

type CreateEmployeeBreakRequest struct {
	Id          uuid.UUID `json:"id" db:"id"`
	EmployeeId  uuid.UUID `json:"employeeId" db:"employee_id"`
	Weekday     int       `json:"weekday" db:"weekday"`
	StartTime   string    `json:"startTime" db:"start_time"`
	EndTime     string    `json:"endTime" db:"end_time"`
	Description string    `json:"description" db:"description"`
	Active      *bool     `json:"active" db:"active"`
}

type UpdateEmployeeBreakRequest struct {
	Id          uuid.UUID `json:"id" db:"id"`
	StartTime   *string   `json:"startTime,omitempty" db:"start_time"`
	EndTime     *string   `json:"endTime,omitempty" db:"end_time"`
	Description *string   `json:"description,omitempty" db:"description"`
	Active      *bool     `json:"active,omitempty" db:"active"`
}
