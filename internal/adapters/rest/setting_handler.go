package rest

import (
	"fmt"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type SettingHandler interface {
	GetPricePerPerson(c *fiber.Ctx) error
	SetPricePerPerson(c *fiber.Ctx) error
}

type settingHandler struct {
	service usecases.SettingUseCase
}

func NewSettingHandler(service usecases.SettingUseCase) SettingHandler {
	return &settingHandler{
		service: service,
	}
}

// GetPricePerPerson
// @Summary Get Price Per Person
// @Description Get price per person in setting.
// @Tags Manage
// @Accept json
// @Produce json
// @Success 200 {object} responses.SettingResponse
// @Router /manage/settings/price-per-person [get]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (s *settingHandler) GetPricePerPerson(c *fiber.Ctx) error {
	price, err := s.service.GetPricePerPerson(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(price)
}

// SetPricePerPerson
// @Summary Set Price Per Person
// @Description Update price per person in setting.
// @Tags Manage
// @Accept json
// @Produce json
// @Param request body requests.EditPricePerPerson true "Edit Price Per Person Request"
// @Success 200 {object} responses.SuccessResponse
// @Router /manage/settings/price-per-person [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (s *settingHandler) SetPricePerPerson(c *fiber.Ctx) error {
	var req requests.EditPricePerPerson
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	if err := s.service.SetPricePerPerson(c.Context(), fmt.Sprintf("%f", req.Price)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Price per person updated",
	})
}
