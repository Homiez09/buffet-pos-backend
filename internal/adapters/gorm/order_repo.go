package gorm

import (
	"context"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/cs471-buffetpos/buffet-pos-backend/domain/requests"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderGormRepository struct {
	DB *gorm.DB
}

func NewOrderGormRepository(db *gorm.DB) *OrderGormRepository {
	return &OrderGormRepository{
		DB: db,
	}
}

func (o *OrderGormRepository) GetOrdersByStatus(ctx context.Context, status string) ([]*models.Order, error) {
	var orders []*models.Order
	result := o.DB.Where("status = ?", status).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (o *OrderGormRepository) GetOrdersByTableID(ctx context.Context, tableID string) ([]*models.Order, error) {
	var orders []*models.Order
	result := o.DB.Where("table_id = ?", tableID).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

func (o *OrderGormRepository) UpdateOrderStatus(ctx context.Context, orderID string, status string) error {
	result := o.DB.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status)
	return result.Error
}

func (o *OrderGormRepository) CreateOrder(ctx context.Context, order *requests.UserAddOrderRequest, tableID string) (*models.Order, error) {
	tableIDParse, err := uuid.Parse(tableID)
	if err != nil {
		return nil, err
	}

	id := uuid.New()

	newOrder := &models.Order{
		ID:      id,
		TableID: tableIDParse,
		Status:  models.PREPARING,
	}

	result := o.DB.Create(newOrder)
	if result.Error != nil {
		return nil, result.Error
	}

	return newOrder, nil
}
