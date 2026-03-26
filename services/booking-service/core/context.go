package core

import (
	"booking-service/common"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type WebContext struct {
	echo.Context
	*AppContext
}

type Context interface {
	GetDb() *sqlx.DB
	GetConfig() *common.Config
}

type AppContext struct {
	config *common.Config
	db     *sqlx.DB
}

func (ac *AppContext) GetDb() *sqlx.DB {
	return ac.db
}

func (ac *AppContext) GetConfig() *common.Config {
	return ac.config
}

func CreateCtx(ctx *AppContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &WebContext{Context: c, AppContext: ctx}
			return next(cc)
		}
	}
}

func (c *WebContext) Success(data ...interface{}) error {

	if data != nil {
		return c.JSON(http.StatusOK, data[0])
	}

	return c.NoContent(http.StatusOK)
}

func (c *WebContext) InternalError(msg string) error {
	return c.JSON(http.StatusInternalServerError, msg)
}

func (c *WebContext) Forbidden(msg string) error {
	return c.JSON(http.StatusForbidden, msg)
}

func (c *WebContext) Unauthorized(msg string) error {
	return c.JSON(http.StatusUnauthorized, msg)
}

func (c *WebContext) BadRequest(msg string) error {
	return c.JSON(http.StatusBadRequest, msg)
}

func (c *WebContext) NotFound(msg string) error {
	return c.JSON(http.StatusNotFound, msg)
}
