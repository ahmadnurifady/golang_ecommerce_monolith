package domain

import "time"

type Order struct {
	OrderId         string    `json:"order_id"`
	UserId          string    `json:"user_id"`
	OrderStatus     string    `json:"order_status"`
	TotalAmount     float64   `json:"total_amount"`
	ShippingAddress string    `json:"shipping_address"`
	PaymentMethod   string    `json:"payment_method"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type OrderItem struct {
	OrderItemId string    `json:"order_item_id"`
	OrderId     string    `json:"order_id"`
	ProductId   string    `json:"product_id"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
