package rest

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/responses"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/usecases"
	"github.com/gofiber/fiber/v2"
)

type OrderHandler interface {
	GetOrdersByStatus(c *fiber.Ctx) error
	GetOrdersByTableID(c *fiber.Ctx) error
	UpdateOrderStatus(c *fiber.Ctx) error
	CreateOrder(c *fiber.Ctx) error
	CustomerGetOrderHistory(c *fiber.Ctx) error
}

type orderHandler struct {
	service usecases.OrderUseCase
}

func NewOrderHandler(service usecases.OrderUseCase) OrderHandler {
	return &orderHandler{
		service: service,
	}
}

// Get Orders By Status
// @Summary Get Orders By Status
// @Description Get orders by status.
// @Tags Manage
// @Accept json
// @Produce json
// @Success 200 {array} responses.OrderDetail
// @Router /manage/orders/status/:status [get]
// @Param status path string true "Order Status"
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *orderHandler) GetOrdersByStatus(c *fiber.Ctx) error {
	status := c.Params("status")
	orders, err := h.service.GetOrdersByStatus(c.Context(), status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(orders)
}

// Get Orders By Table ID
// @Summary Get Orders By Table ID
// @Description Get orders by table ID.
// @Tags Manage
// @Accept json
// @Produce json
// @Success 200 {array} responses.OrderDetail
// @Router /manage/orders/table/:tableID [get]
// @Param tableID path string true "Table ID"
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *orderHandler) GetOrdersByTableID(c *fiber.Ctx) error {
	tableID := c.Params("tableID")
	orders, err := h.service.GetOrdersByTableID(c.Context(), tableID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(orders)
}

// Update Order Status
// @Summary Update Order Status
// @Description Update order status.
// @Tags Manage
// @Accept json
// @Produce json
// @Param request body requests.UpdateOrderStatusRequest true "Update Order Status Request"
// @Success 200 {object} responses.SuccessResponse
// @Router /manage/orders/status [put]
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
func (h *orderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	var req requests.UpdateOrderStatusRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	if err := h.service.UpdateOrderStatus(c.Context(), req.OrderID, req.Status); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Order status updated successfully",
	})
}

// Add Order
// @Summary Add Order
// @Description Add order to table.
// @Tags Customer
// @Accept json
// @Produce json
// @Param request body requests.UserAddOrderRequest true "User Add Order Request"
// @Success 200 {object} responses.SuccessResponse
// @Router /customer/orders [post]
// @param AccessCode header string true "Access Code"
func (h *orderHandler) CreateOrder(c *fiber.Ctx) error {
	var req requests.UserAddOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	claims, ok := c.Locals("table").(*responses.TableDetail)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	tableID := claims.ID

	if err := h.service.CreateOrder(c.Context(), &req, tableID.String()); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"message": "Order created successfully",
	})
}

// Customer Get Order History
// @Summary Get Order History
// @Description Get order history for table.
// @Tags Customer
// @Accept json
// @Produce json
// @Success 200 {array} responses.OrderDetail
// @Router /customer/orders/history [get]
// @param AccessCode header string true "Access Code"
func (h *orderHandler) CustomerGetOrderHistory(c *fiber.Ctx) error {
	claims, ok := c.Locals("table").(*responses.TableDetail)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	tableID := claims.ID
	orders, err := h.service.GetOrderHistory(c.Context(), tableID.String())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(orders)
}
