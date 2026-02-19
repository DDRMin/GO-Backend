package orders

type OrderItem struct {
	ProductID int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

type CreateOrderRequest struct {
	UserID int64       `json:"user_id"`
	Items  []OrderItem `json:"items"`
}

type CreateOrderResponse struct {
	OrderID int64 `json:"order_id"`
}
