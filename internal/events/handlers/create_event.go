package handlers

import (
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"Zondrics/internal/events/service"
	"github.com/gofiber/fiber/v2"
)

func CreateEventHandler(c *fiber.Ctx) error {
	data := new(models.EventInput)

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

	if data.OrganizationUserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'organization_user_id' cannot be empty",
		})
	}

	if data.ResourceID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'resource_id' cannot be empty",
		})
	}

	if data.EventTime == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'event_time' cannot be empty",
		})
	}

	if data.Capacity == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'capacity' cannot be empty",
		})
	}

	formatedTime, err := service.TimeFormatter(data.EventTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Wrong time format",
		})
	}

	checkedTime, err := service.CheckTime(formatedTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Time should be in the future",
		})
	}

	var resource models.Resource
	if err := database.DB.First(&resource, data.ResourceID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Resource not found",
		})
	}

	var organization models.Organization
	if err := database.DB.First(&organization, data.OrganizationID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization not found",
		})
	}

	var organizationUser models.Organization
	if err := database.DB.First(&organizationUser, data.OrganizationUserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization user not found",
		})
	}

	record := models.Event{
		OrganizationID:     data.OrganizationID,
		OrganizationUserID: data.OrganizationUserID,
		ResourceID:         data.ResourceID,
		EventTime:          checkedTime,
		Capacity:           data.Capacity,
	}

	if err := database.DB.Create(&record).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create event record",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event record was created successfully",
	})
}
