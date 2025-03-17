package handlers

import (
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateOrganizationUserHandler(c *fiber.Ctx) error {
	data := new(models.OrganizationUserRegistrationInput)

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

	if data.Login == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Login' cannot be empty",
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

	if data.Phone == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Phone' cannot be empty",
		})
	}

	if database.ValidatePhone(data.Phone) != true {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wrong phone number format",
		})
	}

	if data.Type == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Type' cannot be empty",
		})
	}

	if data.Type != "admin" && data.Type != "employee" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wrong type for 'Type'",
		})
	}

	if data.OrganizationID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'OrganizationID' cannot be empty",
		})
	}

	var existingOrganization models.Organization
	if err := database.DB.First(&existingOrganization, "id = ?", data.OrganizationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Organization doesn't exist",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query database",
		})
	}

	var existingUser models.User
	if err := database.DB.Where("login = ?", data.Login).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Login is already taken",
		})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query database",
		})
	}

	hash := database.Hash(data.Login, data.Password)

	newUser := models.User{
		Login:   data.Login,
		Hash:    hash,
		Name:    data.Name,
		Surname: data.Surname,
		Phone:   data.Phone,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	newOrganizationUser := models.OrganizationUser{
		Type:           data.Type,
		UserID:         newUser.ID,
		OrganizationID: data.OrganizationID,
	}

	if err := database.DB.Create(&newOrganizationUser).Error; err != nil {
		database.DB.Delete(&newUser)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create organization user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Organization user created successfully",
	})
}
