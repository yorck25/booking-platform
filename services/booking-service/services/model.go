package services

import (
	"time"

	"github.com/google/uuid"
)

type Service struct {
	Id              uuid.UUID `json:"id" db:"id"`
	BarberId        uuid.UUID `json:"barberId" db:"barber_id"`
	InternalName    string    `json:"internalName" db:"internal_name"`
	DisplayName     string    `json:"displayName" db:"display_name"`
	Description     string    `json:"description" db:"description"`
	DurationMinutes int       `json:"durationMinutes" db:"duration_minutes"`
	PriceCents      int       `json:"priceCents" db:"price_cents"`
	Active          bool      `json:"active" db:"active"`
	Deleted         bool      `json:"deleted" db:"deleted"`
	SortOrder       int       `json:"sortOrder" db:"sort_order"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}

type CreateServiceRequest struct {
	Id              uuid.UUID `json:"id" db:"id"`
	BarberId        uuid.UUID `json:"barberId" db:"barber_id"`
	InternalName    string    `json:"internalName" db:"internal_name"`
	DisplayName     string    `json:"displayName" db:"display_name"`
	Description     string    `json:"description" db:"description"`
	DurationMinutes int       `json:"durationMinutes" db:"duration_minutes"`
	PriceCents      int       `json:"priceCents" db:"price_cents"`
	Active          *bool     `json:"active" db:"active"`
	SortOrder       int       `json:"sortOrder" db:"sort_order"`
}

type UpdateServiceRequest struct {
	Id              uuid.UUID `json:"id" db:"id"`
	BarberId        uuid.UUID `json:"barberId" db:"barber_id"`
	InternalName    *string   `json:"internalName,omitempty" db:"internal_name"`
	DisplayName     *string   `json:"displayName,omitempty" db:"display_name"`
	Description     *string   `json:"description,omitempty" db:"description"`
	DurationMinutes *int      `json:"durationMinutes,omitempty" db:"duration_minutes"`
	PriceCents      *int      `json:"priceCents,omitempty" db:"price_cents"`
	Active          *bool     `json:"active,omitempty" db:"active"`
	SortOrder       *int      `json:"sortOrder,omitempty" db:"sort_order"`
}

type GetServiceRequest struct {
	Id       uuid.UUID `json:"id" param:"id"`
	BarberId uuid.UUID `json:"barberId" query:"barberId"`
}

type ListServicesRequest struct {
	BarberId uuid.UUID `json:"barberId" query:"barberId"`
	Active   *bool     `json:"active" query:"active"`
}

type DeleteServiceRequest struct {
	Id       uuid.UUID `json:"id" param:"id"`
	BarberId uuid.UUID `json:"barberId" db:"barber_id"`
}
