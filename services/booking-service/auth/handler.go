package auth

import (
	"booking-service/barbers"
	"booking-service/common"
	"booking-service/core"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"

	"time"
)

func HandleVerifyAuth(ctx *core.WebContext) error {
	barberId := ctx.Request().Header.Get("barberID")
	if barberId == "" {
		return ctx.BadRequest("barberID required")
	}

	barberStore := barbers.NewStore(ctx)

	barber, err := barberStore.GetBarberByID(barberId)
	if err != nil {
		return ctx.InternalError("Fail to get Barber")
	}

	if !barber.Active {
		return ctx.NotFound("Barber is not active")
	}

	token, err := createVerifyToken(barber.ID, ctx.GetConfig())
	if err != nil {
		return ctx.InternalError("Fail to create verify token")
	}

	return ctx.Success(token)
}

func createVerifyToken(barberId uuid.UUID, config *common.Config) (string, error) {
	claims := &VerifyClaims{
		BarberId: barberId,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecretKey)
}

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
