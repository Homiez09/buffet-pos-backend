package responses

import (
	"time"

	"github.com/google/uuid"
)

type BaseInvoice struct {
	ID           uuid.UUID  `json:"id"`
	PeopleAmount int        `json:"peopleAmount"`
	TotalPrice   float64    `json:"totalPrice"`
	IsPaid       bool       `json:"isPaid"`
	TableID      *uuid.UUID `json:"tableId"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

type InvoiceDetail struct {
	BaseInvoice
}
