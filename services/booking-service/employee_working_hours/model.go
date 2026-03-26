package employee_working_hours

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeWorkingHour struct {
	Id         uuid.UUID `json:"id" db:"id"`
	EmployeeId uuid.UUID `json:"employeeId" db:"employee_id"`
	Weekday    int       `json:"weekday" db:"weekday"`
	StartTime  string    `json:"startTime" db:"start_time"`
	EndTime    string    `json:"endTime" db:"end_time"`
	IsClosed   bool      `json:"isClosed" db:"is_closed"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}

type CreateEmployeeWorkingHourRequest struct {
	Id         uuid.UUID `json:"id" db:"id"`
	EmployeeId uuid.UUID `json:"employeeId" db:"employee_id"`
	Weekday    int       `json:"weekday" db:"weekday"`
	StartTime  string    `json:"startTime" db:"start_time"`
	EndTime    string    `json:"endTime" db:"end_time"`
	IsClosed   bool      `json:"isClosed" db:"is_closed"`
}

type UpdateEmployeeWorkingHourRequest struct {
	Id        uuid.UUID `json:"id" db:"id"`
	StartTime *string   `json:"startTime,omitempty" db:"start_time"`
	EndTime   *string   `json:"endTime,omitempty" db:"end_time"`
	IsClosed  *bool     `json:"isClosed,omitempty" db:"is_closed"`
}

type ListEmployeeWorkingHoursRequest struct {
	EmployeeId uuid.UUID `json:"employeeId" param:"employeeId"`
}

type GetEmployeeWorkingHourRequest struct {
	EmployeeId uuid.UUID `json:"employeeId" param:"employeeId"`
	Id         uuid.UUID `json:"id" param:"id"`
}
