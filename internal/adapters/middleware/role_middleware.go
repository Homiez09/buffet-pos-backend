package middleware

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func RoleMiddleware(allowedRoles ...models.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(jwt.MapClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Unauthorized",
			})
		}

		role, ok := claims["role"]
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"message": "Forbidden: No role assigned",
			})
		}
		role = models.Role(role.(string))

		for _, allowedRole := range allowedRoles {
			if role == allowedRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Forbidden: Insufficient role",
		})
	}
}
