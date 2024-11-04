package repositories

import (
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
)

type OrderRepository interface {
	GetOrdersByStatus(status string) ([]*models.Order, error)
	GetOrdersByTableID(tableID string) ([]*models.Order, error)
	UpdateOrderStatus(orderID string, status string) error
	CreateOrder(order *requests.UserAddOrderRequest) (*models.Order, error)
}
