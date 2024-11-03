package models

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	OrderID  uuid.UUID `json:"orderID" gorm:"type:uuid;foreignKey:OrderID"`
	MenuID   uuid.UUID `json:"menuID" gorm:"type:uuid;foreignKey:MenuID"`
	Quantity int       `json:"quantity" gorm:"type:integer"`
}
