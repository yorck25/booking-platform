package barber_employees

import (
	"booking-service/core"
	"database/sql"
	"time"

	"github.com/google/uuid"
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
		INSERT INTO barber_employees (
			id,
			barber_id,
			username,
			first_name,
			last_name,
			display_name,
			internal_name,
			email,
			phone,
			password_hash,
			role,
			failed_attempts,
			active,
			deleted,
			created_at,
			updated_at
		) VALUES (
			:id,
			:barber_id,
			:username,
			:first_name,
			:last_name,
			:display_name,
			:internal_name,
			:email,
			:phone,
			:password_hash,
			:role,
			:failed_attempts,
			:active,
			:deleted,
			:created_at,
			:updated_at
		)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()

	params := map[string]any{
		"id":              cbur.Id,
		"barber_id":       cbur.BarberId,
		"username":        cbur.Username,
		"first_name":      cbur.FirstName,
		"last_name":       cbur.LastName,
		"display_name":    cbur.DisplayName,
		"internal_name":   cbur.InternalName,
		"email":           cbur.Email,
		"phone":           cbur.Phone,
		"password_hash":   cbur.PasswordHash,
		"role":            cbur.Role,
		"failed_attempts": 0,
		"active":          true,
		"deleted":         false,
		"created_at":      now,
		"updated_at":      now,
	}

	_, err = stmt.Exec(params)
	return err
}

func (s *Store) GetBarberUserByUsername(barberId uuid.UUID, username string) (*BarberUser, error) {
	var user BarberUser

	err := s.db.Get(&user, `
		SELECT
			id,
			barber_id,
			username,
			first_name,
			last_name,
			display_name,
			internal_name,
			email,
			phone,
			password_hash,
			role,
			last_login,
			failed_attempts,
			active,
			deleted,
			created_at,
			updated_at
		FROM barber_employees
		WHERE barber_id = $1
		  AND username = $2
		  AND deleted = false
		LIMIT 1
	`, barberId, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (s *Store) UpdateLastLogin(userId uuid.UUID) error {
	_, err := s.db.Exec(`
		UPDATE barber_employees
		SET
			last_login = NOW(),
			failed_attempts = 0,
			updated_at = NOW()
		WHERE id = $1
	`, userId)

	return err
}

func (s *Store) IncrementFailedAttempts(userId uuid.UUID) error {
	_, err := s.db.Exec(`
		UPDATE barber_employees
		SET
			failed_attempts = failed_attempts + 1,
			updated_at = NOW()
		WHERE id = $1
	`, userId)

	return err
}
