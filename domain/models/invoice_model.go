package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID           uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid"`
	PeopleAmount int        `json:"peopleAmount" gorm:"type:integer"`
	TotalPrice   float64    `json:"totalPrice" gorm:"type:decimal"`
	IsPaid       bool       `json:"isPaid" gorm:"type:boolean"`
	TableID      *uuid.UUID `json:"tableId" gorm:"type:uuid;foreignKey:TableID"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt
}
