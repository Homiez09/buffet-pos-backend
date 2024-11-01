package responses

import (
	"time"

	"github.com/google/uuid"
)

type BaseTable struct {
	ID          uuid.UUID `json:"id"`
	TableName   string    `json:"tableName"`
	IsAvailable bool      `json:"isAvailable"`
	QRCode      *string   `json:"qrcode,omitempty"`
	AccessCode  *string   `json:"accessCode,omitempty"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type TableDetail struct {
	BaseTable
	// without created_at and updatedAt
}
