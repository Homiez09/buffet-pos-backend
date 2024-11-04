package main

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/bootstrap"
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/gorm"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/middleware"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/adapters/rest"
	"github.com/cs471-buffetpos/buffet-pos-backend/internal/infrastructure/cloudinary"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	cld := cloudinary.NewCloudinaryStorageService(cfg)

	settingRepo := gorm.NewSettingGormRepository(db)
	settingService := usecases.NewSettingService(settingRepo, cfg)
	settingHandler := rest.NewSettingHandler(settingService)

	userRepo := gorm.NewUserGormRepository(db)
	userService := usecases.NewUserService(userRepo, cfg)
	userHandler := rest.NewUserHandler(userService)

	invoiceRepo := gorm.NewInvoiceGormRepository(db)
	tableRepo := gorm.NewTableGormRepository(db)

	invoiceService := usecases.NewInvoiceService(invoiceRepo, tableRepo, cfg)
	tableService := usecases.NewTableService(tableRepo, invoiceRepo, settingRepo, cfg)

	tableHandler := rest.NewTableHandler(tableService)
	invoiceHandler := rest.NewInvoiceHandler(invoiceService)

	categoryRepo := gorm.NewCategoryGormRepository(db)
	categoryService := usecases.NewCategoryService(categoryRepo, cfg)
	categoryHandler := rest.NewCategoryHandler(categoryService)

	menuRepo := gorm.NewMenuGormRepository(db)
	menuService := usecases.NewMenuService(menuRepo, cfg, cld)
	menuHandler := rest.NewMenuHandler(menuService)

	orderRepo := gorm.NewOrderGormRepository(db)
	orderItemRepo := gorm.NewOrderItemGormRepository(db)
	orderService := usecases.NewOrderService(orderRepo, orderItemRepo, menuRepo, cfg)
	orderHandler := rest.NewOrderHandler(orderService)

	app.Use(cors.New())

	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("BuffetPOS is running 🎉")
	})

	auth := app.Group("/auth")
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	customer := app.Group("/customer", middleware.CustomerMiddleware(cfg, tableService))
	customer.Get("/menus", menuHandler.CustomerFindAll)
	customer.Post("/orders", orderHandler.CreateOrder)
	customer.Get("/orders/history", orderHandler.CustomerGetOrderHistory)
	customer.Get("/categories", categoryHandler.FindAllCategories)

	manage := app.Group("/manage", middleware.AuthMiddleware(cfg), middleware.RoleMiddleware(models.Employee, models.Manager))
	manage.Get("/tables", tableHandler.FindAllTables)
	manage.Get("/tables/:id", tableHandler.FindTableByID)
	manage.Post("/tables", tableHandler.AddTable)
	manage.Put("/tables", tableHandler.Edit)
	manage.Delete("/tables/:id", tableHandler.Delete)
	manage.Post("/tables/assign", tableHandler.AssignTable)

	manage.Get("/invoices/paid", invoiceHandler.GetAllPaidInvoices)
	manage.Get("/invoices/unpaid", invoiceHandler.GetAllUnpaidInvoices)
	manage.Put("/invoices/set-paid", invoiceHandler.SetInvoicePaid)
	manage.Delete("/invoices/cancel", invoiceHandler.CancelInvoice)

	manage.Get("/categories", categoryHandler.FindAllCategories)
	manage.Get("/categories/:id", categoryHandler.FindCategoryByID)
	manage.Post("/categories", categoryHandler.AddCategory)
	manage.Delete("/categories/:id", categoryHandler.DeleteCategory)

	manage.Get("/menus", menuHandler.FindAll)
	manage.Get("/menus/:id", menuHandler.FindByID)
	manage.Post("/menus", menuHandler.Create)
	manage.Put("/menus", menuHandler.Edit)
	manage.Delete("/menus/:id", menuHandler.Delete)

	manage.Get("/orders", orderHandler.GetOrdersByStatus)
	manage.Get("/orders/tables", orderHandler.GetOrdersByTableID)
	manage.Put("/orders/status", orderHandler.UpdateOrderStatus)

	manage.Get("/settings/price-per-person", settingHandler.GetPricePerPerson)
	manage.Put("/settings/price-per-person", settingHandler.SetPricePerPerson)

	app.Listen(":3000")
}
