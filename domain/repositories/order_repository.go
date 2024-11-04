package repositories

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
)

type OrderRepository interface {
	GetOrdersByStatus(ctx context.Context, status string) ([]models.Order, error)
	GetOrdersByTableID(ctx context.Context, tableID string) ([]models.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID string, status string) error
	CreateOrder(ctx context.Context, order *requests.UserAddOrderRequest, tableID string) (*models.Order, error)
	GetOrderHistory(ctx context.Context, tableID string) ([]models.Order, error)
}
