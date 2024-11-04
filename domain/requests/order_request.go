package requests

type UserAddOrderRequest struct {
	OrderItems []OrderItemRequest `json:"order_items" validate:"required"`
}

type UpdateOrderStatusRequest struct {
	OrderID string `json:"table_id" validate:"required"`
	Status  string `json:"status" validate:"required"`
}
