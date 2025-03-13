package registration

import (
	"Zondrics/internal/database"
	"github.com/gofiber/fiber/v2"
)

func SetupRegistrationRoutes(app *fiber.App) {
	app.Get("/register", Register)

	app.Post("/create_user", func(c *fiber.Ctx) error {
		data := new(database.User)

		if err := c.BodyParser(data); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if data.Name == "" || data.Hash == "" || data.Phone == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Name and hash cannot be empty",
			})
		}

		if database.ValidatePhone(data.Phone) == false {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid phone number",
			})
		}

		var existingUser database.User
		if err := database.DB.Where("Phone = ?", data.Phone).First(&existingUser).Error; err == nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Phone number already exists",
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
