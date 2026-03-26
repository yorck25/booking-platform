package barber_employees

import (
	"booking-service/auth"
	"booking-service/core"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func HandleCreateBarberUser(ctx *core.WebContext) error {
	var cbur CreateBarberUserRequest

	err := ctx.Bind(&cbur)
	if err != nil {
		return ctx.BadRequest("invalid request")
	}

	if cbur.BarberId == uuid.Nil {
		return ctx.BadRequest("barberId is required")
	}

	if strings.TrimSpace(cbur.Username) == "" {
		return ctx.BadRequest("username is required")
	}

	if strings.TrimSpace(cbur.FirstName) == "" {
		return ctx.BadRequest("firstName is required")
	}

	if strings.TrimSpace(cbur.LastName) == "" {
		return ctx.BadRequest("lastName is required")
	}

	if strings.TrimSpace(cbur.Password) == "" {
		return ctx.BadRequest("password is required")
	}

	if strings.TrimSpace(cbur.DisplayName) == "" {
		cbur.DisplayName = strings.TrimSpace(cbur.FirstName + " " + cbur.LastName)
	}

	if strings.TrimSpace(cbur.InternalName) == "" {
		cbur.InternalName = strings.ToLower(strings.ReplaceAll(cbur.DisplayName, " ", "-"))
	}

	if strings.TrimSpace(cbur.Role) == "" {
		cbur.Role = "user"
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
		return ctx.InternalError("internal error, fail to create barber user")
	}

	token, err := auth.CreateToken(cbur.Id, cbur.Role, cbur.BarberId, ctx.GetConfig())
	if err != nil {
		return ctx.InternalError("internal error, fail to create token")
	}

	user := BarberUser{
		Id:           cbur.Id,
		BarberId:     cbur.BarberId,
		Username:     cbur.Username,
		FirstName:    cbur.FirstName,
		LastName:     cbur.LastName,
		DisplayName:  cbur.DisplayName,
		InternalName: cbur.InternalName,
		Email:        cbur.Email,
		Phone:        cbur.Phone,
		Role:         cbur.Role,
		Active:       true,
		Deleted:      false,
	}

	return ctx.Success(CreateBarberUserResponse{
		Token: token,
		User:  user,
	})
}

func HandleLoginBarberUser(ctx *core.WebContext) error {
	var req LoginBarberUserRequest

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.BadRequest("invalid request")
	}

	if req.BarberId == uuid.Nil {
		return ctx.BadRequest("barberId is required")
	}

	if strings.TrimSpace(req.Username) == "" {
		return ctx.BadRequest("username is required")
	}

	if strings.TrimSpace(req.Password) == "" {
		return ctx.BadRequest("password is required")
	}

	store := NewStore(ctx)

	user, err := store.GetBarberUserByUsername(req.BarberId, req.Username)
	if err != nil {
		return ctx.InternalError("internal error")
	}

	if user == nil {
		return ctx.Unauthorized("invalid username or password")
	}

	if !user.Active || user.Deleted {
		return ctx.Unauthorized("user is inactive")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		_ = store.IncrementFailedAttempts(user.Id)
		return ctx.Unauthorized("invalid username or password")
	}

	err = store.UpdateLastLogin(user.Id)
	if err != nil {
		return ctx.InternalError("internal error")
	}

	token, err := auth.CreateToken(user.Id, user.Role, user.BarberId, ctx.GetConfig())
	if err != nil {
		return ctx.InternalError("internal error, fail to create token")
	}

	return ctx.Success(LoginBarberUserResponse{
		Token: token,
	})
}
