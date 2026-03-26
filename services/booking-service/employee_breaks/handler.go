package employee_breaks

import (
	"booking-service/core"
	"strings"

	"github.com/google/uuid"
)

func HandleCreateEmployeeBreak(ctx *core.WebContext) error {
	employeeId, err := uuid.Parse(strings.TrimSpace(ctx.Param("employeeId")))
	if err != nil {
		return ctx.BadRequest("invalid employeeId")
	}

	var req CreateEmployeeBreakRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest("invalid request")
	}

	req.Id = uuid.New()
	req.EmployeeId = employeeId
	req.StartTime = strings.TrimSpace(req.StartTime)
	req.EndTime = strings.TrimSpace(req.EndTime)
	req.Description = strings.TrimSpace(req.Description)

	if req.Weekday < 0 || req.Weekday > 6 {
		return ctx.BadRequest("weekday must be between 0 and 6")
	}

	if req.StartTime == "" {
		return ctx.BadRequest("startTime is required")
	}

	if req.EndTime == "" {
		return ctx.BadRequest("endTime is required")
	}

	store := NewStore(ctx)

	item, err := store.Create(req)
	if err != nil {
		return ctx.InternalError("internal error, fail to create break")
	}

	return ctx.Success(item)
}

func HandleListEmployeeBreaks(ctx *core.WebContext) error {
	employeeId, err := uuid.Parse(strings.TrimSpace(ctx.Param("employeeId")))
	if err != nil {
		return ctx.BadRequest("invalid employeeId")
	}

	store := NewStore(ctx)

	items, err := store.ListByEmployeeId(employeeId)
	if err != nil {
		return ctx.InternalError("internal error, fail to list breaks")
	}

	return ctx.Success(items)
}

func HandleGetEmployeeBreak(ctx *core.WebContext) error {
	employeeId, err := uuid.Parse(strings.TrimSpace(ctx.Param("employeeId")))
	if err != nil {
		return ctx.BadRequest("invalid employeeId")
	}

	id, err := uuid.Parse(strings.TrimSpace(ctx.Param("id")))
	if err != nil {
		return ctx.BadRequest("invalid id")
	}

	store := NewStore(ctx)

	item, err := store.GetById(employeeId, id)
	if err != nil {
		return ctx.InternalError("internal error, fail to get break")
	}
	if item == nil {
		return ctx.NotFound("break not found")
	}

	return ctx.Success(item)
}

func HandleUpdateEmployeeBreak(ctx *core.WebContext) error {
	employeeId, err := uuid.Parse(strings.TrimSpace(ctx.Param("employeeId")))
	if err != nil {
		return ctx.BadRequest("invalid employeeId")
	}

	id, err := uuid.Parse(strings.TrimSpace(ctx.Param("id")))
	if err != nil {
		return ctx.BadRequest("invalid id")
	}

	var req UpdateEmployeeBreakRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest("invalid request")
	}

	req.Id = id

	if req.StartTime != nil {
		value := strings.TrimSpace(*req.StartTime)
		if value == "" {
			return ctx.BadRequest("startTime must not be empty")
		}
		req.StartTime = &value
	}

	if req.EndTime != nil {
		value := strings.TrimSpace(*req.EndTime)
		if value == "" {
			return ctx.BadRequest("endTime must not be empty")
		}
		req.EndTime = &value
	}

	if req.Description != nil {
		value := strings.TrimSpace(*req.Description)
		req.Description = &value
	}

	store := NewStore(ctx)

	item, err := store.Update(employeeId, req)
	if err != nil {
		return ctx.InternalError("internal error, fail to update break")
	}
	if item == nil {
		return ctx.NotFound("break not found")
	}

	return ctx.Success(item)
}

func HandleDeleteEmployeeBreak(ctx *core.WebContext) error {
	employeeId, err := uuid.Parse(strings.TrimSpace(ctx.Param("employeeId")))
	if err != nil {
		return ctx.BadRequest("invalid employeeId")
	}

	id, err := uuid.Parse(strings.TrimSpace(ctx.Param("id")))
	if err != nil {
		return ctx.BadRequest("invalid id")
	}

	store := NewStore(ctx)

	if err := store.Delete(employeeId, id); err != nil {
		return ctx.InternalError("internal error, fail to delete break")
	}

	return ctx.Success("break deleted successfully")
}
