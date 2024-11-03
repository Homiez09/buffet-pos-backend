package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Table struct {
	ID          uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid"`
	TableName   string     `json:"tableName" gorm:"type:varchar(255)"`
	IsAvailable bool       `json:"isAvailable" gorm:"type:boolean;default:true"`
	QRCode      *string    `json:"qrcode" gorm:"type:varchar(255)"`
	AccessCode  *string    `json:"accessCode" gorm:"type:varchar(255)"`
	EntryAt     *time.Time `json:"entryAt"`
	Invoices    []Invoice  `json:"invoices" gorm:"foreignKey:TableID"`
	Orders      []Order    `json:"orders" gorm:"foreignKey:TableID"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt
}
