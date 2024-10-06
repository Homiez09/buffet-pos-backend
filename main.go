package main

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/bootstrap"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/gorm"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/rest"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/cs471-buffetpos/buffet-pos-backend/docs"
)

// @title BuffetPOS API
// @description This is the BuffetPOS API documentation.
// @version 1.0
func main() {
	app := fiber.New()
	cfg := configs.NewConfig()

	db := bootstrap.NewDB(cfg)

	userRepo := gorm.NewUserGormRepository(db)
	userService := usecases.NewUserService(userRepo, cfg)
	userHandler := rest.NewUserHandler(userService)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BuffetPOS is running 🎉")
	})

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	app.Listen(":3000")
}
