package models

import "github.com/google/uuid"

type BuffetPack struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name     string    `json:"name" gorm:"type:varchar(255)"`
	Price    float64   `json:"price" gorm:"type:decimal"`
	Menus    []Menu    `json:"menus" gorm:"many2many:buffet_pack_menus"`
	Invoices []Invoice `json:"invoices" gorm:"foreignKey:BuffetPackID"`
}
