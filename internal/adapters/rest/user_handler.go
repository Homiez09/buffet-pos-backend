package rest

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
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
