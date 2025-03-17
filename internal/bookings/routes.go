package bookings

import (
	"Zondrics/internal/bookings/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupBookingRoute(app *fiber.App) {
	app.Post("/new_booking", handlers.CreateBookingHandler)
}
