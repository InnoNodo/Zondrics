package registration

import (
	"Zondrics/internal/database"
	"github.com/gofiber/fiber/v2"
)

func SetupRegistrationRoutes(app *fiber.App) {
	app.Get("/register", Register)

	app.Post("/create_user", func(c *fiber.Ctx) error {
		data := new(User)

		if err := c.BodyParser(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if data.Username == "" || data.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Username and password cannot be empty",
			})
		}

		if err := database.DB.Create(&data).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create user",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User created successfully",
			"user":    data,
		})
	})
}
