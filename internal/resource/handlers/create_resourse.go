package handlers

import (
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateResource(c *fiber.Ctx) error {
	data := new(models.ResourceInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input data",
		})
	}

	if data.OrganizationID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'organization_id' cannot be empty",
		})
	}

	if data.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'name' cannot be empty",
		})
	}

	if data.Type == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'type' cannot be empty",
		})
	}

	if data.Location == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'location' cannot be empty",
		})
	}

	if data.Capacity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'capacity' cannot be empty",
		})
	}

	if data.Status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'status' cannot be empty",
		})
	}
	if data.Status != "active" && data.Status != "inactive" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wrong type in field 'status'",
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

	var resource models.Resource
	if err := database.DB.First(&resource, "name = ?", data.Name).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Resource already exists",
		})
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query database",
		})
	}

	record := models.Resource{
		OrganizationID: data.OrganizationID,
		Name:           data.Name,
		Type:           data.Type,
		Location:       data.Location,
		Capacity:       data.Capacity,
		Status:         data.Status,
	}

	if err := database.DB.Create(&record).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create resource",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Resource was created successfully",
	})

}
