package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Menu struct {
	ID          uuid.UUID   `json:"id" gorm:"primaryKey;type:uuid"`
	Name        string      `json:"name" gorm:"type:varchar(255)"`
	Description *string     `json:"description" gorm:"type:text"`
	CategoryID  *uuid.UUID  `json:"categoryId" gorm:"type:uuid"`
	ImageURL    *string     `json:"imageUrl" gorm:"type:varchar(255)"`
	IsAvailable bool        `json:"isAvailable" gorm:"type:boolean"`
	OrderItems  []OrderItem `json:"orderItems" gorm:"foreignKey:MenuID"`
	DeletedAt   gorm.DeletedAt
}
