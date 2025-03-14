package authorization

import (
	"Zondrics/internal/authorization/routings"
	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/auth", routings.LoginHandler)
}
