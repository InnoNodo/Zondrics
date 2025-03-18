package main

import (
	"Zondrics/internal/authorization"
	"Zondrics/internal/bookings"
	"Zondrics/internal/registration"
	"Zondrics/internal/resource"
	"github.com/gofiber/fiber/v2"
)

func CombineRoutes(app *fiber.App) {
	authorization.SetupAuthRoutes(app)
	registration.SetupRegistrationRoutes(app)
	resource.SetupResourceRoutes(app)
	bookings.SetupBookingRoute(app)
}
