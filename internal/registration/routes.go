package registration

import (
	"Zondrics/internal/registration/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRegistrationRoutes(app *fiber.App) {

	//authorized := app.Group("/api", middleware.JWTMiddleware)

	app.Post("/create_user", handlers.CreateUserHandler)

	app.Post("/create_organization", handlers.CreateOrganizationHandler)

	app.Post("/create_organization_user", handlers.CreateOrganizationUserHandler)
}
