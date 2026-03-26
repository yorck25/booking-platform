package services

import (
	"booking-service/core"
	"strings"

	"github.com/google/uuid"
)

func HandleCreateService(ctx *core.WebContext) error {
	var req CreateServiceRequest

	err := ctx.Bind(&req)
	if err != nil {
		return ctx.BadRequest("invalid request")
	}

	if req.BarberId == uuid.Nil {
		return ctx.BadRequest("barberId is required")
	}

	req.DisplayName = strings.TrimSpace(req.DisplayName)
	req.InternalName = strings.TrimSpace(req.InternalName)
	req.Description = strings.TrimSpace(req.Description)

	if req.DisplayName == "" {
		return ctx.BadRequest("displayName is required")
	}

	if req.InternalName == "" {
		req.InternalName = strings.ToLower(strings.ReplaceAll(req.DisplayName, " ", "-"))
	}

	if req.DurationMinutes <= 0 {
		return ctx.BadRequest("durationMinutes must be greater than 0")
	}

	if req.PriceCents < 0 {
		return ctx.BadRequest("priceCents must be greater or equal to 0")
	}

	if req.SortOrder < 0 {
		return ctx.BadRequest("sortOrder must be greater or equal to 0")
	}

	req.Id = uuid.New()

	store := NewStore(ctx)

	service, err := store.CreateService(req)
	if err != nil {
		return ctx.InternalError("internal error, fail to create service")
	}

	return ctx.Success(service)
}

func HandleGetService(ctx *core.WebContext) error {
	serviceIdParam := strings.TrimSpace(ctx.Param("id"))
	if serviceIdParam == "" {
		return ctx.BadRequest("service id is required")
	}

	serviceId, err := uuid.Parse(serviceIdParam)
	if err != nil {
		return ctx.BadRequest("invalid service id")
	}

	barberIdParam := strings.TrimSpace(ctx.QueryParam("barberId"))
	if barberIdParam == "" {
		return ctx.BadRequest("barberId is required")
	}

	barberId, err := uuid.Parse(barberIdParam)
	if err != nil {
		return ctx.BadRequest("invalid barberId")
	}

	store := NewStore(ctx)

	service, err := store.GetServiceById(barberId, serviceId)
	if err != nil {
		return ctx.InternalError("internal error, fail to get service")
	}
	if service == nil {
		return ctx.NotFound("service not found")
	}

	return ctx.Success(service)
}

func HandleListServices(ctx *core.WebContext) error {
	barberIdParam := strings.TrimSpace(ctx.QueryParam("barberId"))
	if barberIdParam == "" {
		return ctx.BadRequest("barberId is required")
	}

	barberId, err := uuid.Parse(barberIdParam)
	if err != nil {
		return ctx.BadRequest("invalid barberId")
	}

	var activeFilter *bool
	activeParam := strings.TrimSpace(ctx.QueryParam("active"))
	if activeParam != "" {
		switch activeParam {
		case "true":
			v := true
			activeFilter = &v
		case "false":
			v := false
			activeFilter = &v
		default:
			return ctx.BadRequest("active must be true or false")
		}
	}

	store := NewStore(ctx)

	services, err := store.ListServices(barberId, activeFilter)
	if err != nil {
		return ctx.InternalError("internal error, fail to list services")
	}

	return ctx.Success(services)
}

func HandleUpdateService(ctx *core.WebContext) error {
	serviceIdParam := strings.TrimSpace(ctx.Param("id"))
	if serviceIdParam == "" {
		return ctx.BadRequest("service id is required")
	}

	serviceId, err := uuid.Parse(serviceIdParam)
	if err != nil {
		return ctx.BadRequest("invalid service id")
	}

	var req UpdateServiceRequest
	err = ctx.Bind(&req)
	if err != nil {
		return ctx.BadRequest("invalid request")
	}

	if req.BarberId == uuid.Nil {
		return ctx.BadRequest("barberId is required")
	}

	req.Id = serviceId

	if req.DisplayName != nil {
		value := strings.TrimSpace(*req.DisplayName)
		if value == "" {
			return ctx.BadRequest("displayName must not be empty")
		}
		req.DisplayName = &value
	}

	if req.InternalName != nil {
		value := strings.TrimSpace(*req.InternalName)
		if value == "" {
			return ctx.BadRequest("internalName must not be empty")
		}
		req.InternalName = &value
	}

	if req.Description != nil {
		value := strings.TrimSpace(*req.Description)
		req.Description = &value
	}

	if req.DurationMinutes != nil && *req.DurationMinutes <= 0 {
		return ctx.BadRequest("durationMinutes must be greater than 0")
	}

	if req.PriceCents != nil && *req.PriceCents < 0 {
		return ctx.BadRequest("priceCents must be greater or equal to 0")
	}

	if req.SortOrder != nil && *req.SortOrder < 0 {
		return ctx.BadRequest("sortOrder must be greater or equal to 0")
	}

	store := NewStore(ctx)

	service, err := store.UpdateService(req)
	if err != nil {
		return ctx.InternalError("internal error, fail to update service")
	}
	if service == nil {
		return ctx.NotFound("service not found")
	}

	return ctx.Success(service)
}

func HandleDeleteService(ctx *core.WebContext) error {
	serviceIdParam := strings.TrimSpace(ctx.Param("id"))
	if serviceIdParam == "" {
		return ctx.BadRequest("service id is required")
	}

	serviceId, err := uuid.Parse(serviceIdParam)
	if err != nil {
		return ctx.BadRequest("invalid service id")
	}

	barberIdParam := strings.TrimSpace(ctx.QueryParam("barberId"))
	if barberIdParam == "" {
		return ctx.BadRequest("barberId is required")
	}

	barberId, err := uuid.Parse(barberIdParam)
	if err != nil {
		return ctx.BadRequest("invalid barberId")
	}

	store := NewStore(ctx)

	err = store.DeleteService(barberId, serviceId)
	if err != nil {
		return ctx.InternalError("internal error, fail to delete service")
	}

	return ctx.Success("service deleted successfully")
}
