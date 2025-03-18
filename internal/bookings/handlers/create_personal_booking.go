package handlers

import (
	"Zondrics/internal/bookings/service"
	"Zondrics/internal/database"
	"Zondrics/internal/database/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

func CreatePersonalBookingHandler(c *fiber.Ctx) error {
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
			"error": "Field 'user_id' cannot be empty",
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

	if data.StartTime == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Event Date cannot be empty",
		})
	}

	formatedTime, err := service.TimeFormatter(data.StartTime)
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

		if data.OrganizationUserID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Field 'organization_user_id' cannot be empty",
			})
		}

		var organizationUser models.OrganizationUser
		if err := database.DB.First(&organizationUser, "id = ?", data.OrganizationUserID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Trainer not found",
			})
		}

		if data.UserID == organizationUser.UserID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Trainer and user cannot be the same",
			})
		}

		if data.OrganizationID == organizationUser.OrganizationID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Trainer not correlated to organization",
			})
		}

		timeSlot := models.Timeslot{
			UserID:             &data.UserID,
			OrganizationID:     data.OrganizationID,
			OrganizationUserID: &data.OrganizationUserID,
			StartTime:          checkedTime,
			EndTime:            checkedTime.Add(time.Hour),
		}

		record := models.TrainingBooking{
			UserID:             data.UserID,
			OrganizationID:     data.OrganizationID,
			OrganizationUserID: data.OrganizationUserID,
			ResourceID:         data.ResourceID,
			Timeslot:           timeSlot,
			Age:                data.Age,
		}
		if err := database.DB.Create(&record).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create personal events record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Personal training booking created successfully",
		})

	case "haircut":
		if data.Price <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Price must be provided and greater than 0 for haircut bookings",
			})
		}

		var organizationUser models.OrganizationUser
		if err := database.DB.First(&organizationUser, data.OrganizationUserID).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Hairdresser not found",
			})
		}

		if data.UserID == organizationUser.UserID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Hairdresser and user cannot be the same",
			})
		}

		if data.OrganizationID == organizationUser.OrganizationID {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Hairdresser not correlated to organization",
			})
		}

		timeSlot := models.Timeslot{
			UserID:             &data.UserID,
			OrganizationID:     data.OrganizationID,
			OrganizationUserID: &data.OrganizationUserID,
			StartTime:          checkedTime,
			EndTime:            checkedTime.Add(time.Hour),
		}

		record := models.HaircutBooking{
			UserID:             data.UserID,
			OrganizationID:     data.OrganizationID,
			OrganizationUserID: data.OrganizationUserID,
			ResourceID:         data.ResourceID,
			Timeslot:           timeSlot,
			Price:              data.Price,
		}
		if err := database.DB.Create(&record).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create haircut record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Haircut booking created successfully",
		})

	case "restaurant":
		if data.Guests <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Guests must be provided and greater than 0 for restaurant bookings",
			})
		}

		timeSlot := models.Timeslot{
			UserID:             &data.UserID,
			OrganizationID:     data.OrganizationID,
			OrganizationUserID: nil,
			StartTime:          checkedTime,
			EndTime:            checkedTime.Add(time.Hour),
		}

		record := models.RestaurantBooking{
			UserID:         data.UserID,
			OrganizationID: data.OrganizationID,
			TableNumber:    data.TableNumber,
			ResourceID:     data.ResourceID,
			Timeslot:       timeSlot,
			Guests:         data.Guests,
		}
		if err := database.DB.Create(&record).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to create restaurant booking record",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Restaurant booking created successfully",
		})

	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Unsupported activity type",
		})
	}
}
