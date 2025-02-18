package authorization

import "github.com/gofiber/fiber/v2"

func SetupAuthRoutes(app *fiber.App) {
	app.Get("/auth", Authorization)
}
