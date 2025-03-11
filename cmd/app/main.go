package main

import (
	"Zondrics/internal/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {

	if err := database.InitDatabase(); err != nil {
		fmt.Println("Error while database initialization:", err)
		return
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	CombineRoutes(app)

	app.Listen(":3000")
}
