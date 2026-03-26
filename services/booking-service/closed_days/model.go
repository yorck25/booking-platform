package closed_days

import (
	"time"

	"github.com/google/uuid"
)

type ClosedDay struct {
	Id         uuid.UUID `json:"id" db:"id"`
	BarberId   uuid.UUID `json:"barberId" db:"barber_id"`
	ClosedDate string    `json:"closedDate" db:"closed_date"`
	Reason     string    `json:"reason" db:"reason"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

type CreateClosedDayRequest struct {
	Id         uuid.UUID `json:"id" db:"id"`
	BarberId   uuid.UUID `json:"barberId" db:"barber_id"`
	ClosedDate string    `json:"closedDate" db:"closed_date"`
	Reason     string    `json:"reason" db:"reason"`
}

type UpdateClosedDayRequest struct {
	Id         uuid.UUID `json:"id" db:"id"`
	ClosedDate *string   `json:"closedDate,omitempty" db:"closed_date"`
	Reason     *string   `json:"reason,omitempty" db:"reason"`
}
