package registration

import (
	"Zondrics/internal/registration/routings"
	"github.com/gofiber/fiber/v2"
)

func SetupRegistrationRoutes(app *fiber.App) {

	//authorized := app.Group("/api", middleware.JWTMiddleware)
	app.Post("/create_user", routings.CreateUserHandler)
}
