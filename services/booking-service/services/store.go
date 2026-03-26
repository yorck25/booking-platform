package services

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

func (s *Store) CreateService(req CreateServiceRequest) (*Service, error) {
	now := time.Now()

	active := true
	if req.Active != nil {
		active = *req.Active
	}

	service := &Service{
		Id:              req.Id,
		BarberId:        req.BarberId,
		InternalName:    req.InternalName,
		DisplayName:     req.DisplayName,
		Description:     req.Description,
		DurationMinutes: req.DurationMinutes,
		PriceCents:      req.PriceCents,
		Active:          active,
		Deleted:         false,
		SortOrder:       req.SortOrder,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	stmt, err := s.db.PrepareNamed(`
		INSERT INTO services (
			id,
			barber_id,
			internal_name,
			display_name,
			description,
			duration_minutes,
			price_cents,
			active,
			deleted,
			sort_order,
			created_at,
			updated_at
		) VALUES (
			:id,
			:barber_id,
			:internal_name,
			:display_name,
			:description,
			:duration_minutes,
			:price_cents,
			:active,
			:deleted,
			:sort_order,
			:created_at,
			:updated_at
		)
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(service)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s *Store) GetServiceById(barberId, serviceId uuid.UUID) (*Service, error) {
	var service Service

	err := s.db.Get(&service, `
		SELECT
			id,
			barber_id,
			internal_name,
			display_name,
			description,
			duration_minutes,
			price_cents,
			active,
			deleted,
			sort_order,
			created_at,
			updated_at
		FROM services
		WHERE id = $1
		  AND barber_id = $2
		  AND deleted = false
		LIMIT 1
	`, serviceId, barberId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &service, nil
}

func (s *Store) ListServices(barberId uuid.UUID, active *bool) ([]*Service, error) {
	services := make([]*Service, 0)

	query := `
		SELECT
			id,
			barber_id,
			internal_name,
			display_name,
			description,
			duration_minutes,
			price_cents,
			active,
			deleted,
			sort_order,
			created_at,
			updated_at
		FROM services
		WHERE barber_id = $1
		  AND deleted = false
	`
	args := []any{barberId}

	if active != nil {
		query += ` AND active = $2`
		args = append(args, *active)
	}

	query += `
		ORDER BY sort_order ASC, display_name ASC
	`

	err := s.db.Select(&services, query, args...)
	if err != nil {
		return nil, err
	}

	return services, nil
}

func (s *Store) UpdateService(req UpdateServiceRequest) (*Service, error) {
	current, err := s.GetServiceById(req.BarberId, req.Id)
	if err != nil {
		return nil, err
	}
	if current == nil {
		return nil, nil
	}

	if req.InternalName != nil {
		current.InternalName = *req.InternalName
	}
	if req.DisplayName != nil {
		current.DisplayName = *req.DisplayName
	}
	if req.Description != nil {
		current.Description = *req.Description
	}
	if req.DurationMinutes != nil {
		current.DurationMinutes = *req.DurationMinutes
	}
	if req.PriceCents != nil {
		current.PriceCents = *req.PriceCents
	}
	if req.Active != nil {
		current.Active = *req.Active
	}
	if req.SortOrder != nil {
		current.SortOrder = *req.SortOrder
	}

	current.UpdatedAt = time.Now()

	stmt, err := s.db.PrepareNamed(`
		UPDATE services
		SET
			internal_name = :internal_name,
			display_name = :display_name,
			description = :description,
			duration_minutes = :duration_minutes,
			price_cents = :price_cents,
			active = :active,
			sort_order = :sort_order,
			updated_at = :updated_at
		WHERE id = :id
		  AND barber_id = :barber_id
		  AND deleted = false
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

func (s *Store) DeleteService(barberId, serviceId uuid.UUID) error {
	_, err := s.db.Exec(`
		UPDATE services
		SET
			deleted = true,
			active = false,
			updated_at = NOW()
		WHERE id = $1
		  AND barber_id = $2
		  AND deleted = false
	`, serviceId, barberId)

	return err
}
