package handlers

import (
	"Zondrics/internal/bookings/service"
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"github.com/gofiber/fiber/v2"
)

func CreateBookingHandler(c *fiber.Ctx) error {
	UserID, err := service.GetUserIDFromJWT(c)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	bookingInput := new(models.BookingInput)
	if err := c.BodyParser(bookingInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	var organization models.Organization
	if err := database.DB.First(&organization, bookingInput.OrganizationID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Organization not found",
		})
	}

	switch organization.Activity {
	case "sport":
		if bookingInput.Age <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Age must be provided and greater than 0 for sport bookings",
			})
		}
		record := models.TrainingBooking{
			UserID:             UserID,
			OrganizationID:     bookingInput.OrganizationID,
			OrganizationUserID: bookingInput.OrganizationUserID,
			Duration:           bookingInput.Duration,
			EventDate:          bookingInput.EventDate,
			EventTime:          bookingInput.EventTime,
			Age:                bookingInput.Age,
		}
		if err := database.DB.Create(&record).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create personal training record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Personal training record created successfully",
		})

	case "haircut":
		if bookingInput.Price <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Price must be provided and greater than 0 for haircut bookings",
			})
		}
		record := models.HaircutBooking{
			UserID:             UserID,
			OrganizationID:     bookingInput.OrganizationID,
			OrganizationUserID: bookingInput.OrganizationUserID,
			EventDate:          bookingInput.EventDate,
			EventTime:          bookingInput.EventTime,
			Price:              bookingInput.Price,
		}
		if err := database.DB.Create(&record).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create haircut record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Haircut record created successfully",
		})

	case "restaurant":
		if bookingInput.Guests <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Guests must be provided and greater than 0 for restaurant bookings",
			})
		}
		record := models.RestaurantBooking{
			UserID:         UserID,
			OrganizationID: bookingInput.OrganizationID,
			TableNumber:    bookingInput.TableNumber,
			EventDate:      bookingInput.EventDate,
			EventTime:      bookingInput.EventTime,
			Guests:         bookingInput.Guests,
		}
		if err := database.DB.Create(&record).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create restaurant booking record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Restaurant booking record created successfully",
		})

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unsupported activity type",
		})
	}
}
