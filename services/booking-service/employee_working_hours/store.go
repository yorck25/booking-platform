package employee_working_hours

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

func (s *Store) Create(req CreateEmployeeWorkingHourRequest) (*EmployeeWorkingHour, error) {
	now := time.Now().UTC()

	item := &EmployeeWorkingHour{
		Id:         req.Id,
		EmployeeId: req.EmployeeId,
		Weekday:    req.Weekday,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		IsClosed:   req.IsClosed,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	stmt, err := s.db.PrepareNamed(`
		INSERT INTO employee_working_hours (
			id,
			employee_id,
			weekday,
			start_time,
			end_time,
			is_closed,
			created_at,
			updated_at
		) VALUES (
			:id,
			:employee_id,
			:weekday,
			:start_time,
			:end_time,
			:is_closed,
			:created_at,
			:updated_at
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

	return item, nil
}

func (s *Store) ListByEmployeeId(employeeId uuid.UUID) ([]*EmployeeWorkingHour, error) {
	items := make([]*EmployeeWorkingHour, 0)

	err := s.db.Select(&items, `
		SELECT
			id,
			employee_id,
			weekday,
			start_time,
			end_time,
			is_closed,
			created_at,
			updated_at
		FROM employee_working_hours
		WHERE employee_id = $1
		ORDER BY weekday ASC, start_time ASC
	`, employeeId)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s *Store) GetById(employeeId, id uuid.UUID) (*EmployeeWorkingHour, error) {
	var item EmployeeWorkingHour

	err := s.db.Get(&item, `
		SELECT
			id,
			employee_id,
			weekday,
			start_time,
			end_time,
			is_closed,
			created_at,
			updated_at
		FROM employee_working_hours
		WHERE employee_id = $1
		  AND id = $2
		LIMIT 1
	`, employeeId, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &item, nil
}

func (s *Store) Update(employeeId uuid.UUID, req UpdateEmployeeWorkingHourRequest) (*EmployeeWorkingHour, error) {
	current, err := s.GetById(employeeId, req.Id)
	if err != nil {
		return nil, err
	}
	if current == nil {
		return nil, nil
	}

	if req.StartTime != nil {
		current.StartTime = *req.StartTime
	}
	if req.EndTime != nil {
		current.EndTime = *req.EndTime
	}
	if req.IsClosed != nil {
		current.IsClosed = *req.IsClosed
	}

	current.UpdatedAt = time.Now().UTC()

	stmt, err := s.db.PrepareNamed(`
		UPDATE employee_working_hours
		SET
			start_time = :start_time,
			end_time = :end_time,
			is_closed = :is_closed,
			updated_at = :updated_at
		WHERE employee_id = :employee_id
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

	return current, nil
}

func (s *Store) Delete(employeeId, id uuid.UUID) error {
	_, err := s.db.Exec(`
		DELETE FROM employee_working_hours
		WHERE employee_id = $1
		  AND id = $2
	`, employeeId, id)

	return err
}
