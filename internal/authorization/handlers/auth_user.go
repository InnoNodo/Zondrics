package handlers

import (
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func LoginHandler(c *fiber.Ctx) error {
	data := new(models.AuthInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.Login == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Login' cannot be empty",
		})
	}

	if data.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Password' cannot be empty",
		})
	}

	var user models.User
	if err := database.DB.Where("Login = ?", data.Login).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "User not found",
			})
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to query database",
			})
		}
	}

	hash := database.Hash(data.Login, data.Password)

	if user.Hash != hash {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Incorrect password",
		})
	}

	var organizationUser models.OrganizationUser
	err := database.DB.Where("user_id = ?", user.ID).First(&organizationUser).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			token, err := database.CreateTokenForUser(user)
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error": "Failed to generate token",
				})
			}
			c.Set("Authorization", "Bearer "+token)

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "User logged in successfully",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query organization user",
		})
	}

	token, err := database.CreateTokenForOrganizationUser(user, organizationUser.OrganizationID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to generate token",
		})
	}

	c.Set("Authorization", "Bearer "+token)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Organization user logged in successfully",
	})
}
