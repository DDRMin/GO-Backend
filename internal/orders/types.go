package orders

type orderItem struct {
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
}

type createOrderRequest struct {
	UserID    string      `json:"user_id"`
	OrderItems []orderItem `json:"order_items"`
}