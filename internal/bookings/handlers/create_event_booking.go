package handlers

import (
	"Zondrics/internal/bookings/service"
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"github.com/gofiber/fiber/v2"
)

func CreateEventBookingHandler(c *fiber.Ctx) error {
	UserID, err := service.GetUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	data := new(models.BookingInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.ResourceID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Resource ID is required",
		})
	}

	var organization models.Organization
	if err := database.DB.First(&organization, data.OrganizationID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization not found",
		})
	}

	var resource models.Resource
	if err := database.DB.First(&resource, data.ResourceID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Resource not found",
		})
	}

	if data.EventTime == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Event Date cannot be empty",
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

	if data.Age <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Age must be provided and greater than 0 for sport bookings",
		})
	}

	if data.Duration <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Duration must be provided and greater than 0",
		})
	}

	var organizationUser models.Organization
	if err := database.DB.First(&organizationUser, data.OrganizationUserID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Trainer not found",
		})
	}

	record := models.TrainingBooking{
		UserID:             UserID,
		OrganizationID:     data.OrganizationID,
		OrganizationUserID: data.OrganizationUserID,
		ResourceID:         data.ResourceID,
		Duration:           data.Duration,
		EventTime:          checkedTime,
		Age:                data.Age,
	}

	if err := database.DB.Create(&record).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create events booking",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Event booking created successfully",
	})
}
