package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
)

func GetUserIDFromJWT(c *fiber.Ctx) (uint, error) {
	user := c.Locals("user").(*jwt.Token) // Получаем токен из локального контекста
	claims := user.Claims.(jwt.MapClaims)

	// Конвертируем userID в uint
	userID, err := strconv.ParseUint(claims["userID"].(string), 10, 32)
	if err != nil {
		return 0, err
	}

	return uint(userID), nil
}
