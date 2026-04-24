package core

import (
	"booking-service/common"
	"fmt"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type App struct {
	*echo.Echo
	Ctx *AppContext
}

func InitApp() (*App, error) {
	ctx := &AppContext{}

	config, err := LoadConfig()
	if err != nil {
		return nil, err
	}

	ctx.config = config

	fmt.Println("Start connecting to Turso/libSQL!")
	fmt.Println(ctx.config.ConnectionStr)

	dbURL := ctx.config.ConnectionStr

	if ctx.config.TursoToken != "" {
		dbURL = fmt.Sprintf("%s?authToken=%s", ctx.config.ConnectionStr, ctx.config.TursoToken)
	}

	db, err := sqlx.Connect("libsql", dbURL)
	if err != nil {
		return nil, err
	}

	ctx.db = db

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("cannot reach the database: %w", err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		return nil, fmt.Errorf("could not enable foreign keys: %w", err)
	}

	fmt.Println("✅ Successfully connected to Turso/libSQL!")

	e := echo.New()
	e.Debug = true
	e.Logger.SetLevel(log.DEBUG)
	e.Use(CreateCtx(ctx))

	return &App{Echo: e, Ctx: ctx}, nil
}

type HandlerFunc func(*WebContext) error

func (f HandlerFunc) Handle(ctx *WebContext) error {
	return f(ctx)
}

func wrapHandler(h HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.(*WebContext)
		return h.Handle(ctx)
	}
}

func (a *App) Group(prefix string, m ...echo.MiddlewareFunc) *Group {
	g := a.Echo.Group(prefix, m...)
	return &Group{Group: g}
}

type Group struct {
	*echo.Group
}

func (g *Group) GET(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.Add(http.MethodGet, path, wrapHandler(h), m...)
}

func (g *Group) POST(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.Add(http.MethodPost, path, wrapHandler(h), m...)
}

func (g *Group) PUT(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.Add(http.MethodPut, path, wrapHandler(h), m...)
}

func (g *Group) DELETE(path string, h HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route {
	return g.Add(http.MethodDelete, path, wrapHandler(h), m...)
}

func logOutboundIP() {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		fmt.Println("Error getting IP:", err)
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println("Private IP:", localAddr.IP.String())
	publicIP, err := common.GetURLBody("https://api.ipify.org")
	if err != nil {
		fmt.Println("Error getting public IP:", err)
		return
	}

	fmt.Println("Public IP:", publicIP)
}
