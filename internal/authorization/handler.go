package authorization

import "github.com/gofiber/fiber/v2"

func Authorization(c *fiber.Ctx) error {
	return c.SendString("You authorized!")
}
