package middleware

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/configs"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

func CustomerMiddleware(cfg *configs.Config, tableService usecases.TableUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("AccessCode")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		table, err := tableService.FindByAccessCode(c.Context(), authHeader)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if table == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		c.Locals("table", table)

		return c.Next()
	}
}
