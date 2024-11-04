package responses

import (
	"time"

	"github.com/cs471-buffetpos/buffet-pos-backend/domain/models"
	"github.com/google/uuid"
)

type BaseOrder struct {
	ID        uuid.UUID          `json:"id"`
	TableID   uuid.UUID          `json:"tableId"`
	Status    models.OrderStatus `json:"status"`
	OrderItem []models.OrderItem `json:"orderItem"`
	CreatedAt time.Time          `json:"createdAt"`
	UpdatedAt time.Time          `json:"updatedAt"`
}

type OrderDetail struct {
	BaseOrder
}
