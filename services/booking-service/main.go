package main

import (
	"booking-service/barber_employees"
	"booking-service/barbers"
	"booking-service/bookings"
	"booking-service/closed_days"
	"booking-service/core"
	"booking-service/employee_breaks"
	"booking-service/employee_working_hours"
	"booking-service/services"

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
	mainRoot.GET("/barber/availability", barbers.HandleGetReservationSlots)

	// -- Barber-Users --
	mainRoot.POST("/barber/users/register", barber_employees.HandleCreateBarberUser)
	mainRoot.POST("/barber/users/login", barber_employees.HandleLoginBarberUser)

	// -- Services --
	mainRoot.POST("/barber/services", services.HandleCreateService)
	mainRoot.GET("/barber/services/list", services.HandleListServices)
	mainRoot.GET("/barber/services/:id", services.HandleGetService)
	mainRoot.PUT("/barber/services/:id", services.HandleUpdateService)
	mainRoot.DELETE("/barber/services/:id", services.HandleDeleteService)

	// -- Barber Employees Working Hours --
	mainRoot.POST("/employees/:employeeId/working-hours", employee_working_hours.HandleCreateEmployeeWorkingHour)
	mainRoot.GET("/employees/:employeeId/working-hours", employee_working_hours.HandleListEmployeeWorkingHours)
	mainRoot.GET("/employees/:employeeId/working-hours/:id", employee_working_hours.HandleGetEmployeeWorkingHour)
	mainRoot.PUT("/employees/:employeeId/working-hours/:id", employee_working_hours.HandleUpdateEmployeeWorkingHour)
	mainRoot.DELETE("/employees/:employeeId/working-hours/:id", employee_working_hours.HandleDeleteEmployeeWorkingHour)

	// -- Barber Employees Working Hours --
	mainRoot.POST("/employees/:employeeId/breaks", employee_breaks.HandleCreateEmployeeBreak)
	mainRoot.GET("/employees/:employeeId/breaks", employee_breaks.HandleListEmployeeBreaks)
	mainRoot.GET("/employees/:employeeId/breaks/:id", employee_breaks.HandleGetEmployeeBreak)
	mainRoot.PUT("/employees/:employeeId/breaks/:id", employee_breaks.HandleUpdateEmployeeBreak)
	mainRoot.DELETE("/employees/:employeeId/breaks/:id", employee_breaks.HandleDeleteEmployeeBreak)

	// -- Barber Closed Days --
	mainRoot.POST("/closed-days", closed_days.HandleCreateClosedDay)
	mainRoot.GET("/closed-days", closed_days.HandleListClosedDays)
	mainRoot.GET("/closed-days/:id", closed_days.HandleGetClosedDay)
	mainRoot.PUT("/closed-days/:id", closed_days.HandleUpdateClosedDay)
	mainRoot.DELETE("/closed-days/:id", closed_days.HandleDeleteClosedDay)

	app.Logger.Fatal(app.Start("0.0.0.0:8080"))
}

func defaultUrl(ctx *core.WebContext) error {
	return ctx.Success("This is the backend")
}

func healthUrl(ctx *core.WebContext) error {
	return ctx.Success("OK")
}
