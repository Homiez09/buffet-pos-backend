package responses

import (
	"time"

	"github.com/google/uuid"
)

type FindTableResponse struct {
	ID          uuid.UUID `json:"id"`
	TableName   string    `json:"tableName"`
	IsAvailable bool      `json:"isAvailable"`
	QRCode      *string   `json:"qrcode"`
	AccessCode  *string   `json:"accessCode"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type AddTableResponse struct {
	Message string `json:"message"`
}
