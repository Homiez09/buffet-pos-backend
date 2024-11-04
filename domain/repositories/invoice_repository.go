package repositories

import "context"

type InvoiceRepository interface {
	Create(ctx context.Context, tableID string, totalPrice float64, peopleAmount int) error
	SetPaid(ctx context.Context, invoiceID string) error
}
