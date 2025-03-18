package slots

import (
	"Zondrics/internal/slots/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupSlotsRoutes(app *fiber.App) {
	app.Get("/available_personal_slots", handlers.GetAvailablePersonalBookingsHandler)
	app.Get("/available_event_slots", handlers.GetAvailableEventsHandler)
}
