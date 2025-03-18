package handlers

import (
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"Zondrics/internal/slots/service"
	"github.com/gofiber/fiber/v2"
)

func GetAvailableEventsHandler(c *fiber.Ctx) error {
	data := new(models.EventListInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.OrganizationID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'organization_id' cannot be empty",
		})
	}

	if data.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'user_id' cannot be empty",
		})
	}

	events, err := service.GetAvailableEvents(database.DB, data.OrganizationID, data.UserID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error retrieving events",
		})
	}

	return c.JSON(fiber.Map{
		"events": events,
	})
}
