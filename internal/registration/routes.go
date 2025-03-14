package registration

import (
	"Zondrics/internal/registration/routings"
	"github.com/gofiber/fiber/v2"
)

func SetupRegistrationRoutes(app *fiber.App) {
	app.Post("/create_user", routings.CreateUserHandler)
}
