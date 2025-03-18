package handlers

import (
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"Zondrics/internal/slots/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

func GetAvailablePersonalBookingsHandler(c *fiber.Ctx) error {
	data := new(models.CalendarInput)

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

	var organization *models.Organization
	if err := database.DB.First(&organization, data.OrganizationID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization not found",
		})
	}

	formatedOpeningTime, err := service.FormatTime(organization.OpeningTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid time format for field 'opening_time', expected HH:mm",
		})
	}

	formatedClosingTime, err := service.FormatTime(organization.ClosingTime)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid time format for field 'closing_time', expected HH:mm",
		})
	}

	formatedDay, err := service.FormatDate(data.Day)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid day format, expecting YYYY-MM-DD",
		})
	}

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	if formatedDay.Before(today) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'day' cannot be in the past",
		})
	}

	if formatedOpeningTime.After(formatedClosingTime) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'opening_time' cannot be after field 'closing_time'",
		})
	}

	slots, err := service.GetAvailableSlots(
		database.DB,
		data.OrganizationID,
		&data.OrganizationUserID,
		time.Hour,
		formatedOpeningTime,
		formatedClosingTime,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error while getting available slots",
		})
	}

	return c.JSON(fiber.Map{
		"slots": slots,
	})
}
