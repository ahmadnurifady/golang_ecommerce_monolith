package domain

import "time"

type Payment struct {
	PaymentId     string    `json:"payment_id"`
	OrderId       string    `json:"order_id"`
	Amount        float64   `json:"amount"`
	PaymentStatus string    `json:"payment_status"`
	PaymentDate   time.Time `json:"payment_date"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
