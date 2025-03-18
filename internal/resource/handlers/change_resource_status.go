package handlers

import (
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"github.com/gofiber/fiber/v2"
)

func ChangeResourceStatus(c *fiber.Ctx) error {
	data := new(models.ResourceStatusInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.ResourceID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Resource ID cannot be empty",
		})
	}

	if data.Status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Status cannot be empty",
		})
	}

	if data.Status != "active" && data.Status != "inactive" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Status must be 'active' or 'inactive'",
		})
	}

	var resource models.Resource
	if err := database.DB.First(&resource, data.ResourceID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Resource not found",
		})
	}

	if resource.Status == data.Status {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Resource status '" + data.Status + "' is already '" + data.Status + "'",
		})
	}

	resource.Status = data.Status
	if err := database.DB.Save(&resource).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update resource status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Status was changed successfully",
	})
}
