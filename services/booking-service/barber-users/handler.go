package barber_users

import (
	"booking-service/auth"
	"booking-service/core"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HandleCreateBarberUser(ctx *core.WebContext) error {
	var cbur CreateBarberUserRequest

	err := ctx.Bind(&cbur)
	if err != nil {
		return ctx.BadRequest("invalid request")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(cbur.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.InternalError("internal error")
	}
	cbur.PasswordHash = string(passwordHash)
	cbur.Id = uuid.New()

	store := NewStore(ctx)

	err = store.CreateBarberUser(cbur)
	if err != nil {
		return ctx.InternalError("internal error, fail to create BarberUser")
	}

	token, err := auth.CreateToken(cbur.Id, "user", cbur.BarberId, ctx.GetConfig())
	if err != nil {
		return ctx.InternalError("internal error, fail to create Token")
	}

	return ctx.Success(map[string]any{
		"token": token,
	})
}

func HandleLoginBarberUser(ctx *core.WebContext) error {
	var lbur LoginBarberUserRequest

	store := NewStore(ctx)

	token, err := auth.CreateToken(cbur.Id, "user", cbur.BarberId, ctx.GetConfig())
	if err != nil {
		return ctx.InternalError("internal error, fail to create Token")
	}

	return ctx.Success(map[string]any{
		"token": token,
	})
}
