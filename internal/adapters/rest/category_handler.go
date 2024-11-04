package rest

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/exceptions"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/cs471-buffetpos/buffet-pos-backend/utils"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler interface {
	AddCategory(c *fiber.Ctx) error
	FindAllCategories(c *fiber.Ctx) error
	FindCategoryByID(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

type categoryHandler struct {
	service usecases.CategoryUseCase
}

func NewCategoryHandler(service usecases.CategoryUseCase) CategoryHandler {
	return &categoryHandler{
		service: service,
	}
}

// Add Category
// @Summary Add Category
// @Description Add a new category.
// @Tags Manage
// @Accept json
// @Produce json
// @Param request body requests.AddCategoryRequest true "Add Category request"
// @Success 200 {object} responses.SuccessResponse
// @Router /manage/categories [post]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (ct *categoryHandler) AddCategory(c *fiber.Ctx) error {
	var req *requests.AddCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := utils.ValidateStruct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}

	if err := ct.service.AddCategory(c.Context(), req); err != nil {
		switch err {
		case exceptions.ErrDuplicatedCategoryName:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Category name already exists",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category added successfully",
	})
}

// Find All Categories
// @Summary Find All Categories
// @Description Find all categories.
// @Tags Manage
// @Accept json
// @Produce json
// @Success 200 {array} responses.CategoryDetail
// @Router /manage/categories [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (ct *categoryHandler) FindAllCategories(c *fiber.Ctx) error {
	categories, err := ct.service.FindAllCategories(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(categories)
}

// Find Category By ID
// @Summary Find Category By ID
// @Description Find category by ID.
// @Tags Manage
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} responses.CategoryDetail
// @Router /manage/categories/{id} [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (ct *categoryHandler) FindCategoryByID(c *fiber.Ctx) error {
	id, err := utils.ValidateUUID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}

	category, err := ct.service.FindCategoryByID(c.Context(), *id)
	if err != nil {
		switch err {
		case exceptions.ErrCategoryNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Category not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(category)
}

// Delete Category
// @Summary Delete Category
// @Description Delete category by ID.
// @Tags Manage
// @Accept json
// @Produce json
// @Param id path string true "Category ID"
// @Success 200 {object} responses.SuccessResponse
// @Router /manage/categories/{id} [delete]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (ct *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	id, err := utils.ValidateUUID(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID",
		})
	}
	if err := ct.service.DeleteCategory(c.Context(), *id); err != nil {
		switch err {
		case exceptions.ErrCategoryNotFound:
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Category not found",
			})
		default:
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Category deleted successfully",
	})
}

// Customer Get All Categories
// @Summary Get All Categories
// @Description Get all categories.
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {array} responses.CategoryDetail
// @Router /customer/categories [get]
// @param AccessCode header string true "Access Code"
func (ct *categoryHandler) GetAllCategories(c *fiber.Ctx) error {
	categories, err := ct.service.FindAllCategories(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(categories)
}
