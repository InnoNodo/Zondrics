package resource

import (
	"Zondrics/internal/resource/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupResourceRoutes(app *fiber.App) {
	app.Post("/create_resource", handlers.CreateResource)

	app.Post("/change_resource_status", handlers.ChangeResourceStatus)
}
