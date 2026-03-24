package barber_users

import (
	"time"

	"github.com/google/uuid"
)

type BarberUserLogin struct {
	Id             uuid.UUID `json:"id" db:"id"`
	BarberId       uuid.UUID `json:"barberId" db:"barber_id"`
	Username       string    `json:"username" db:"username"`
	Firstname      string    `json:"firstname" db:"first_name"`
	Lastname       string    `json:"lastname" db:"last_name"`
	Email          string    `json:"email" db:"email"`
	Role           string    `json:"role" db:"role"`
	LastLogin      time.Time `json:"lastLogin" db:"last_login"`
	FailedAttempts int8      `json:"failedAttempts" db:"failed_attempts"`
	Active         bool      `json:"active" db:"active"`
}

type CreateBarberUserRequest struct {
	Id           uuid.UUID `json:"id" db:"id"`
	BarberId     uuid.UUID `json:"barberId" db:"barber_id"`
	Username     string    `json:"username" db:"username"`
	Firstname    string    `json:"firstname" db:"firs_name"`
	Lastname     string    `json:"lastname" db:"last_name"`
	Email        string    `json:"email" db:"email"`
	Password     string    `json:"password" db:"password"`
	PasswordHash string    `json:"passwordHash" db:"password_hash"`
	Role         string    `json:"role" db:"role"`
}

type LoginBarberUserRequest struct {
	BarberId uuid.UUID `json:"barberId" db:"barber_id"`
}

type LoginBarberUserResponse struct {
	Token string `json:"token"`
}
