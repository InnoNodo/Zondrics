package service

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func GetUserIDFromJWT(c *fiber.Ctx) (uint, error) {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	userID, err := strconv.ParseUint(claims["userID"].(string), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(userID), nil
}

func TimeFormatter(input string) (time.Time, error) {
	layout := "2006-01-02T15:04:05"

	parsedTime, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, errors.New("wrong time format")
	}

	return parsedTime, nil
}

func CheckTime(times time.Time) (time.Time, error) {
	if times.Before(time.Now()) || times.Equal(time.Now()) {
		return time.Time{}, errors.New("time must be in the future")
	}

	return times, nil
}
