package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type UserClaims struct {
	UserId   uuid.UUID `json:"userId"`
	Role     string    `json:"role"`
	BarberId uuid.UUID `json:"barberId"`
	jwt.StandardClaims
}
