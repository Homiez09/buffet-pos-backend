package models

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	PREPARING OrderStatus = "preparing"
	SERVED    OrderStatus = "served"
	SUCCEEDED OrderStatus = "succeeded"
	CANCELLED OrderStatus = "cancelled"
)

type Order struct {
	ID         uuid.UUID   `json:"id" gorm:"primaryKey;type:uuid"`
	TableID    uuid.UUID   `json:"tableId" gorm:"type:uuid;foreignKey:TableID"`
	Status     OrderStatus `json:"status" gorm:"type:varchar(50);default:'preparing'"`
	OrderItems []OrderItem `json:"orderItems" gorm:"foreignKey:OrderID"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
}
