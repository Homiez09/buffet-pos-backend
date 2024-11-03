package rest

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type MenuHandler interface {
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
}

type menuHandler struct {
	service usecases.MenuUseCase
}

func NewMenuHandler(service usecases.MenuUseCase) MenuHandler {
	return &menuHandler{
		service: service,
	}
}

func (m *menuHandler) Create(c *fiber.Ctx) error {

	var req *requests.AddMenuRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := m.service.Create(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedMenuName:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Menu name already exist",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Create Menu successfully",
	})
}

func (m *menuHandler) FindAll(c *fiber.Ctx) error {
	res, err := m.service.FindAll(c.Context())
	if err != nil {
		switch err {
		case exceptions.ErrMenuNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Menu not have",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Get All Menus",
		"data":    res,
	})
}

func (m *menuHandler) FindByID(c *fiber.Ctx) error {

	var req *requests.MenuFindByIDRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	res, err := m.service.FindByID(c.Context(), req.ID)
	if err != nil {
		switch err {
		case exceptions.ErrMenuNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Menu not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Menu find successfully",
		"data":    res,
	})
}
