package models

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string     `json:"name" gorm:"type:varchar(255)"`
	Menus     []Menu     `json:"menus" gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}
