package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItemGormRepository struct {
	DB *gorm.DB
}

func NewOrderItemGormRepository(db *gorm.DB) *OrderItemGormRepository {
	return &OrderItemGormRepository{
		DB: db,
	}
}

func (o *OrderItemGormRepository) GetOrderItemsByOrderID(ctx context.Context, orderID string) ([]models.OrderItem, error) {
	var orderItems []models.OrderItem
	result := o.DB.Where("order_id = ?", orderID).Find(&orderItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderItems, nil
}

func (o *OrderItemGormRepository) CreateOrderItem(ctx context.Context, orderItem *requests.OrderItemRequest, orderID string) error {
	newOrderID, err := uuid.Parse(orderID)
	if err != nil {
		return err
	}

	menuID, err := uuid.Parse(orderItem.MenuID)
	if err != nil {
		return err
	}

	id := uuid.New()

	newOrderItem := &models.OrderItem{
		ID:       id,
		OrderID:  newOrderID,
		MenuID:   menuID,
		Quantity: orderItem.Quantity,
	}
	result := o.DB.Create(newOrderItem)
	return result.Error
}
