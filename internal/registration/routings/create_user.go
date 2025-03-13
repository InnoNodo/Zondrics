package routings

import (
	"Zondrics/internal/database"
	"github.com/gofiber/fiber/v2"
)

func CreateUserHandler(c *fiber.Ctx) error {
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

	if !database.ValidatePhone(data.Phone) {
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

	newUser := database.User{Name: data.Name, Hash: data.Hash, Phone: data.Phone}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"user":    data,
	})
}
