package bookings

import (
	"Zondrics/internal/bookings/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupBookingRoutes(app *fiber.App) {
	app.Post("/create_personal_booking", handlers.CreatePersonalBookingHandler)
	app.Post("/create_event_booking", handlers.CreateEventBookingHandler)

}
