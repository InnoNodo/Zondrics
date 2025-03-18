package handlers

import (
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"Zondrics/internal/registration/service"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateOrganizationHandler(c *fiber.Ctx) error {
	data := new(models.OrganizationRegistrationInput)

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

	if data.Activity != "sport" && data.Activity != "restaurant" && data.Activity != "haircut" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Activity' have wrong type of activity",
		})
	}

	if data.Activity == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'Activity' cannot be empty",
		})
	}

	if data.City == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'City' cannot be empty",
		})
	}

	if data.OpeningTime == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'OpeningTime' cannot be empty",
		})
	}

	if data.ClosingTime == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'ClosingTime' cannot be empty",
		})
	}

	//if data.OpeningTime == data.ClosingTime {
	//	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"error": "Field 'opening_time' cannot be equal to field 'closing_time'",
	//	})
	//}

	formatedOpeningTime, err := service.FormatTime(data.OpeningTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid time format for field 'opening_time', expected HH:mm",
		})
	}

	formatedClosingTime, err := service.FormatTime(data.ClosingTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid time format for field 'closing_time', expected HH:mm",
		})
	}

	if formatedOpeningTime == formatedClosingTime {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'opening_time' cannot be equal to field 'closing_time'",
		})
	}

	if formatedOpeningTime.After(formatedClosingTime) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'opening_time' cannot be after field 'closing_time'",
		})
	}

	var existingOrganization models.Organization
	if err := database.DB.Where("Name = ?", data.Name).First(&existingOrganization).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
		} else {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to query database",
			})
		}
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Name is already taken",
		})
	}

	newOrganization := models.Organization{Name: data.Name, City: data.City, Activity: data.Activity}

	if err := database.DB.Create(&newOrganization).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create organization",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Organization created successfully",
	})
}
