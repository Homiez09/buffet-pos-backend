package rest

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	Register(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type userHandler struct {
	service usecases.UserUseCase
}

func NewUserHandler(service usecases.UserUseCase) UserHandler {
	return &userHandler{
		service: service,
	}
}

// Register
// @Summary User register (Employee Role by default)
// @Description Register a new user.
// @Tags User
// @Accept json
// @Produce json
// @Param request body requests.UserRegisterRequest true "User register request"
// @Success 200 {object} responses.UserRegisterResponse
// @Router /auth/register [post]
func (u *userHandler) Register(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UserRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Register user
	if err := u.service.Register(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedEmail:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Email already registered",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User registered successfully",
	})
}

// Login
// @Summary User login.
// @Description Check user credentials and return user data.
// @Tags User
// @Accept json
// @Produce json
// @Param request body requests.UserLoginRequest true "User login request"
// @Success 200 {object} responses.UserLoginResponse
// @Router /auth/login [post]
func (u *userHandler) Login(c *fiber.Ctx) error {
	// Parse request
	var req *requests.UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Login user
	user, err := u.service.Login(c.Context(), req)
	if err != nil {
		switch err {
		case exceptions.ErrLoginFailed:
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Login failed",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(user)
}
