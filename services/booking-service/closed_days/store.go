package closed_days

import (
	"booking-service/core"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB
}

func NewStore(ctx *core.WebContext) *Store {
	return &Store{db: ctx.GetDb()}
}

func (s *Store) Create(req CreateClosedDayRequest) (*ClosedDay, error) {
	item := &ClosedDay{
		Id:         req.Id,
		BarberId:   req.BarberId,
		ClosedDate: req.ClosedDate,
		Reason:     req.Reason,
	}

	stmt, err := s.db.PrepareNamed(`
		INSERT INTO barber_closed_days (
			id,
			barber_id,
			closed_date,
			reason
		) VALUES (
			:id,
			:barber_id,
			:closed_date,
			:reason
		)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(item)
	if err != nil {
		return nil, err
	}

	return s.GetById(req.BarberId, req.Id)
}

func (s *Store) List(barberId uuid.UUID) ([]*ClosedDay, error) {
	items := make([]*ClosedDay, 0)

	err := s.db.Select(&items, `
		SELECT
			id,
			barber_id,
			closed_date,
			reason,
			created_at
		FROM barber_closed_days
		WHERE barber_id = $1
		ORDER BY closed_date ASC
	`, barberId)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Store) GetById(barberId, id uuid.UUID) (*ClosedDay, error) {
	var item ClosedDay

	err := s.db.Get(&item, `
		SELECT
			id,
			barber_id,
			closed_date,
			reason,
			created_at
		FROM barber_closed_days
		WHERE barber_id = $1
		  AND id = $2
		LIMIT 1
	`, barberId, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

func (s *Store) Update(barberId uuid.UUID, req UpdateClosedDayRequest) (*ClosedDay, error) {
	current, err := s.GetById(barberId, req.Id)
	if err != nil {
		return nil, err
	}
	if current == nil {
		return nil, nil
	}

	if req.ClosedDate != nil {
		current.ClosedDate = *req.ClosedDate
	}
	if req.Reason != nil {
		current.Reason = *req.Reason
	}

	stmt, err := s.db.PrepareNamed(`
		UPDATE barber_closed_days
		SET
			closed_date = :closed_date,
			reason = :reason
		WHERE barber_id = :barber_id
		  AND id = :id
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(current)
	if err != nil {
		return nil, err
	}

	return s.GetById(barberId, req.Id)
}

func (s *Store) Delete(barberId, id uuid.UUID) error {
	_, err := s.db.Exec(`
		DELETE FROM barber_closed_days
		WHERE barber_id = $1
		  AND id = $2
	`, barberId, id)

	return err
}
