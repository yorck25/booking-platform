package main

import (
	"booking-service/barbers"
	"booking-service/bookings"
	"booking-service/core"

	"github.com/labstack/echo/v4/middleware"
)

type UserClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	app, err := core.InitApp()

	if err != nil {
		panic(err)
	}

	app.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	mainRoot := app.Group("/api/v1")

	mainRoot.GET("/", defaultUrl)
	mainRoot.GET("/health", healthUrl)

	// -- Bookings --
	mainRoot.POST("/bookings", bookings.HandleCreateBooking)
	mainRoot.PUT("/bookings/cancel", bookings.CancelBooking)

	// -- Barber --
	mainRoot.GET("/barber/reservations-slots", barbers.HandleGetReservationSlots)

	app.Logger.Fatal(app.Start("0.0.0.0:8080"))

}
func defaultUrl(ctx *core.WebContext) error {
	return ctx.Success("This is the backend")
}

func healthUrl(ctx *core.WebContext) error {
	return ctx.Success("OK")
}
