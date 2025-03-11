package main

import (
	"Zondrics/internal/authorization"
	"Zondrics/internal/registration"
	"github.com/gofiber/fiber/v2"
)

func CombineRoutes(app *fiber.App) {
	authorization.SetupAuthRoutes(app)
	registration.SetupRegistrationRoutes(app)
}
