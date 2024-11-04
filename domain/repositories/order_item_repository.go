package repositories

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
)

type OrderItemRepository interface {
	GetOrderItemsByOrderID(orderID string) ([]*models.OrderItem, error)
	CreateOrderItem(orderItem *requests.OrderItemRequest, orderID string) error
}
