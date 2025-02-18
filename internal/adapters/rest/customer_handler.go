package rest

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type CustomerHandler interface {
	Register(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	AddPoint(c *fiber.Ctx) error
	RedeemPoint(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type customerHandler struct {
	service usecases.CustomerUseCase
}

func NewCustomerHandler(service usecases.CustomerUseCase) CustomerHandler {
	return &customerHandler{
		service: service,
	}
}

func (m *customerHandler) FindAll(c *fiber.Ctx) error {
	customers, err := m.service.FindAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(customers)
}

func (m *customerHandler) AddPoint(c *fiber.Ctx) error {
	// Parse request
	var req *requests.CustomerAddPointRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	
	// Add Point
	customer, err := m.service.AddPoint(c.Context(), req);
	if  err != nil {
		switch err {
		case exceptions.ErrCustomerNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Not Found Customer",
			})
		case exceptions.ErrIncorrectPIN:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Incorrect PIN",
			})
		case exceptions.ErrPointLimit:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Point is full",
			})
		case exceptions.ErrInvalidPoint:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid point",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func (m *customerHandler) RedeemPoint(c *fiber.Ctx) error {
	// Parse request
	var req *requests.CustomerRedeemRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Redeem Point
	customer, err := m.service.RedeemPoint(c.Context(), req);
	if  err != nil {
		switch err {
		case exceptions.ErrCustomerNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Not Found Customer",
			})
		case exceptions.ErrIncorrectPIN:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Incorrect PIN",
			})
		case exceptions.ErrNotEnoughPoints:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Not Enough Point",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(customer)
}

func (m *customerHandler) Register(c *fiber.Ctx) error {
	// Parse request
	var req *requests.CustomerRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Validate request
	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	// Register customer
	if err := m.service.Register(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedPhone:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Phone already registered",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Customer registered successfully",
	})
}

func (m *customerHandler) Delete(c *fiber.Ctx) error {
	id, err := utils.ValidateUUID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}

	if err := m.service.DeleteCustomer(c.Context(), *id); err != nil {
		switch err {
		case exceptions.ErrCustomerNotFound:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Customer not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Delete Customer successfully",
	})
}

