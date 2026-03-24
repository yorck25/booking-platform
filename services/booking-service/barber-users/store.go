package barber_users

import (
	"booking-service/core"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(ctx *core.WebContext) *Store {
	return &Store{db: ctx.GetDb()}
}

func (s *Store) CreateBarberUser(cbur CreateBarberUserRequest) error {
	stmt, err := s.db.PrepareNamed(`
		INSERT INTO barber_user_login (
			id,
			barber_id,
			username,
			firs_name,
			last_name,
			email,
			password_hash,
			role,
			failed_attempts,
			active
		) VALUES (
			:id,
			:barber_id,
			:username,
			:firs_name,
			:last_name,
			:email,
			:password_hash,
			:role,
			:failed_attempts,
			:active
		)
	`)
	if err != nil {
		return err
	}

	params := map[string]any{
		"barber_id":       cbur.BarberId,
		"username":        cbur.Username,
		"firs_name":       cbur.Firstname,
		"last_name":       cbur.Lastname,
		"email":           cbur.Email,
		"password_hash":   cbur.PasswordHash,
		"role":            cbur.Role,
		"failed_attempts": 0,
		"active":          true,
	}

	err = stmt.Get(stmt, params)

	return err
}
