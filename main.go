package main

import (
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cs471-buffetpos/buffet-pos-backend/bootstrap"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/gorm"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/middleware"
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
	cld, err := cloudinary.NewFromURL(cfg.Cloudinary_url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cld)

	db := bootstrap.NewDB(cfg)

	userRepo := gorm.NewUserGormRepository(db)
	userService := usecases.NewUserService(userRepo, cfg)
	userHandler := rest.NewUserHandler(userService)

	tableRepo := gorm.NewTableGormRepository(db)
	tableService := usecases.NewTableService(tableRepo, cfg)
	tableHandler := rest.NewTableHandler(tableService)

	categoryRepo := gorm.NewCategoryGormRepository(db)
	categoryService := usecases.NewCategoryService(categoryRepo, cfg)
	categoryHandler := rest.NewCategoryHandler(categoryService)

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BuffetPOS is running 🎉")
	})

	auth := app.Group("/auth")
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	manage := app.Group("/manage", middleware.AuthMiddleware(cfg), middleware.RoleMiddleware(models.Employee, models.Manager))
	manage.Get("/tables", tableHandler.FindAllTables)
	manage.Get("/tables/:id", tableHandler.FindTableByID)
	manage.Post("/tables", tableHandler.AddTable)

	manage.Get("/categories", categoryHandler.FindAllCategories)
	manage.Get("/categories/:id", categoryHandler.FindCategoryByID)
	manage.Post("/categories", categoryHandler.AddCategory)

	app.Listen(":3000")
}
