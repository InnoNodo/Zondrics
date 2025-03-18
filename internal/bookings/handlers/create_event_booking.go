package handlers

import (
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"github.com/gofiber/fiber/v2"
)

func CreateEventBookingHandler(c *fiber.Ctx) error {
	//  Get user_id form JWT-token

	//UserID, err := service.GetUserIDFromJWT(c)
	//if err != nil {
	//	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	//		"error": "Unauthorized",
	//	})
	//}

	data := new(models.BookingInput)

	if err := c.BodyParser(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if data.UserID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'user_id' cannot be empty'",
		})
	}

	if data.EventID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Field 'event_id' cannot be empty'",
		})
	}

	var event *models.Event
	if err := database.DB.First(&event, data.EventID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Event not found",
		})
	}

	var count int64
	err := database.DB.Model(&models.EventBooking{}).
		Where("user_id = ? AND event_id = ?", data.UserID, data.EventID).
		Count(&count).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to query database",
		})
	}

	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Event booking already exists",
		})
	}

	record := models.EventBooking{
		UserID:  data.UserID,
		EventID: data.EventID,
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
