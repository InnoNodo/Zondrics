package events

import (
	"Zondrics/internal/events/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupEventRoutes(app *fiber.App) {
	app.Post("/create_event", handlers.CreateEventHandler)

}
