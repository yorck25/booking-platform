package barber_employees

import (
	"time"

	"github.com/google/uuid"
)

type BarberUser struct {
	Id             uuid.UUID  `json:"id" db:"id"`
	BarberId       uuid.UUID  `json:"barberId" db:"barber_id"`
	Username       string     `json:"username" db:"username"`
	FirstName      string     `json:"firstName" db:"first_name"`
	LastName       string     `json:"lastName" db:"last_name"`
	DisplayName    string     `json:"displayName" db:"display_name"`
	InternalName   string     `json:"internalName" db:"internal_name"`
	Email          string     `json:"email" db:"email"`
	Phone          string     `json:"phone" db:"phone"`
	PasswordHash   string     `json:"-" db:"password_hash"`
	Role           string     `json:"role" db:"role"`
	LastLogin      *time.Time `json:"lastLogin" db:"last_login"`
	FailedAttempts int        `json:"failedAttempts" db:"failed_attempts"`
	Active         bool       `json:"active" db:"active"`
	Deleted        bool       `json:"deleted" db:"deleted"`
	CreatedAt      time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time  `json:"updatedAt" db:"updated_at"`
}

type CreateBarberUserRequest struct {
	Id           uuid.UUID `json:"id" db:"id"`
	BarberId     uuid.UUID `json:"barberId" db:"barber_id"`
	Username     string    `json:"username" db:"username"`
	FirstName    string    `json:"firstName" db:"first_name"`
	LastName     string    `json:"lastName" db:"last_name"`
	DisplayName  string    `json:"displayName" db:"display_name"`
	InternalName string    `json:"internalName" db:"internal_name"`
	Email        string    `json:"email" db:"email"`
	Phone        string    `json:"phone" db:"phone"`
	Password     string    `json:"password" db:"password"`
	PasswordHash string    `json:"passwordHash" db:"password_hash"`
	Role         string    `json:"role" db:"role"`
}

type LoginBarberUserRequest struct {
	BarberId uuid.UUID `json:"barberId" db:"barber_id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
}

type LoginBarberUserResponse struct {
	Token string `json:"token"`
}

type CreateBarberUserResponse struct {
	Token string     `json:"token"`
	User  BarberUser `json:"user"`
}
