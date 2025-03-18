package handlers

import (
	"Zondrics/internal/database/models"
	"github.com/gofiber/fiber/v2"
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
}
