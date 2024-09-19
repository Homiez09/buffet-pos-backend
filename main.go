package main

import (
	"fmt"

	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	cfg := configs.NewConfig()

	fmt.Println("DBHost:", cfg.DBHost)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
