package models

import (
	"time"

	"github.com/google/uuid"
)

type Invoice struct {
	ID           uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	PeopleAmount int       `json:"peopleAmount" gorm:"type:integer"`
	TotalPrice   float64   `json:"totalPrice" gorm:"type:decimal"`
	IsPaid       bool      `json:"isPaid" gorm:"type:boolean"`
	BuffetPackID uuid.UUID `json:"buffetPackId" gorm:"type:uuid;foreignKey:BuffetPackID"`
	TableID      uuid.UUID `json:"tableId" gorm:"type:uuid;foreignKey:TableID"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
