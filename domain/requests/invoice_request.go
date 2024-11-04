package requests

type UpdateInvoiceStatusRequest struct {
	InvoiceID string `json:"invoice_id" validate:"required"`
}
