package closed_days

import (
	"booking-service/core"
	"strings"

	"github.com/google/uuid"
)

func HandleCreateClosedDay(ctx *core.WebContext) error {
	var req CreateClosedDayRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest("invalid request")
	}

	if req.BarberId == uuid.Nil {
		return ctx.BadRequest("barberId is required")
	}

	req.Id = uuid.New()
	req.ClosedDate = strings.TrimSpace(req.ClosedDate)
	req.Reason = strings.TrimSpace(req.Reason)

	if req.ClosedDate == "" {
		return ctx.BadRequest("closedDate is required")
	}

	store := NewStore(ctx)

	item, err := store.Create(req)
	if err != nil {
		return ctx.InternalError("internal error, fail to create closed day")
	}

	return ctx.Success(item)
}

func HandleListClosedDays(ctx *core.WebContext) error {
	barberId, err := uuid.Parse(strings.TrimSpace(ctx.QueryParam("barberId")))
	if err != nil {
		return ctx.BadRequest("invalid barberId")
	}

	store := NewStore(ctx)

	items, err := store.List(barberId)
	if err != nil {
		return ctx.InternalError("internal error, fail to list closed days")
	}

	return ctx.Success(items)
}

func HandleGetClosedDay(ctx *core.WebContext) error {
	id, err := uuid.Parse(strings.TrimSpace(ctx.Param("id")))
	if err != nil {
		return ctx.BadRequest("invalid id")
	}

	barberId, err := uuid.Parse(strings.TrimSpace(ctx.QueryParam("barberId")))
	if err != nil {
		return ctx.BadRequest("invalid barberId")
	}

	store := NewStore(ctx)

	item, err := store.GetById(barberId, id)
	if err != nil {
		return ctx.InternalError("internal error, fail to get closed day")
	}
	if item == nil {
		return ctx.NotFound("closed day not found")
	}

	return ctx.Success(item)
}

func HandleUpdateClosedDay(ctx *core.WebContext) error {
	id, err := uuid.Parse(strings.TrimSpace(ctx.Param("id")))
	if err != nil {
		return ctx.BadRequest("invalid id")
	}

	var req UpdateClosedDayRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.BadRequest("invalid request")
	}

	barberId, err := uuid.Parse(strings.TrimSpace(ctx.QueryParam("barberId")))
	if err != nil {
		return ctx.BadRequest("invalid barberId")
	}

	req.Id = id

	if req.ClosedDate != nil {
		value := strings.TrimSpace(*req.ClosedDate)
		if value == "" {
			return ctx.BadRequest("closedDate must not be empty")
		}
		req.ClosedDate = &value
	}

	if req.Reason != nil {
		value := strings.TrimSpace(*req.Reason)
		req.Reason = &value
	}

	store := NewStore(ctx)

	item, err := store.Update(barberId, req)
	if err != nil {
		return ctx.InternalError("internal error, fail to update closed day")
	}
	if item == nil {
		return ctx.NotFound("closed day not found")
	}

	return ctx.Success(item)
}

func HandleDeleteClosedDay(ctx *core.WebContext) error {
	id, err := uuid.Parse(strings.TrimSpace(ctx.Param("id")))
	if err != nil {
		return ctx.BadRequest("invalid id")
	}

	barberId, err := uuid.Parse(strings.TrimSpace(ctx.QueryParam("barberId")))
	if err != nil {
		return ctx.BadRequest("invalid barberId")
	}

	store := NewStore(ctx)

	if err := store.Delete(barberId, id); err != nil {
		return ctx.InternalError("internal error, fail to delete closed day")
	}

	return ctx.Success("closed day deleted successfully")
}
