package registration

import "github.com/gofiber/fiber/v2"

func SetupRegistrationRoutes(app *fiber.App) {
	app.Get("/register", Register)
}
