package auth

import (
	"booking-service/common"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"time"
)

func CreateToken(userId uuid.UUID, role string, barberId uuid.UUID, config *common.Config) (string, error) {
	claims := &UserClaims{
		UserId:   userId,
		Role:     role,
		BarberId: barberId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(48 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecretKey)
}

func DecodeToken(tokenString string, jwtSecret []byte) (uuid.UUID, error) {
	claims := UserClaims{}

	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return claims.UserId, err
	}

	return claims.UserId, nil
}
