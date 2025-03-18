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

	switch organization.Activity {
	case "sport":
		if data.Age <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Age must be provided and greater than 0 for sport bookings",
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
				"error": "Failed to create personal events record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Personal events record created successfully",
		})

	case "haircut":
		if data.Price <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Price must be provided and greater than 0 for haircut bookings",
			})
		}

		var organizationUser models.Organization
		if err := database.DB.First(&organizationUser, data.OrganizationUserID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Hairdresser not found",
			})
		}

		record := models.HaircutBooking{
			UserID:             UserID,
			OrganizationID:     data.OrganizationID,
			OrganizationUserID: data.OrganizationUserID,
			ResourceID:         data.ResourceID,
			EventTime:          checkedTime,
			Price:              data.Price,
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
		if data.Guests <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Guests must be provided and greater than 0 for restaurant bookings",
			})
		}
		record := models.RestaurantBooking{
			UserID:         UserID,
			OrganizationID: data.OrganizationID,
			TableNumber:    data.TableNumber,
			ResourceID:     data.ResourceID,
			EventTime:      checkedTime,
			Guests:         data.Guests,
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
