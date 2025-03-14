package routings

import (
	"Zondrics/internal/database"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUserHandler(c *fiber.Ctx) error {
	data := new(database.RegistrationInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Name' cannot be empty",
		})
	}

	if data.Surname == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Surname' cannot be empty",
		})
	}

	if data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Password' cannot be empty",
		})
	}

	if !database.ValidatePhone(data.Phone) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid phone number",
		})
	}

	var existingUser database.User
	if err := database.DB.Where("Login = ?", data.Login).First(&existingUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to query database",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Login is already taken",
		})
	}

	hash := database.Hash(data.Login, data.Password)

	//if err != nil {
	//	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	//		"error": "Failed to hash password",
	//	})
	//}

	newUser := database.User{Name: data.Name, Hash: hash, Phone: data.Phone, Surname: data.Surname, Login: data.Login}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
		"hash":    hash,
	})
}
