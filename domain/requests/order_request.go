package requests

type UserAddOrderRequest struct {
	TableID    string             `json:"table_id" validate:"required"`
	OrderItems []OrderItemRequest `json:"order_items" validate:"required"`
}

type UpdateOrderStatusRequest struct {
	OrderID string `json:"table_id" validate:"required"`
	Status  string `json:"status" validate:"required"`
}
