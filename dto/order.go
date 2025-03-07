package dto

import "time"

type CreateOrder struct {
	UserID uint           `json:"user_id"`
	Items  []OrderItemDTO `json:"items"`
}

type UpdateStatusRequest struct {
	Status string `json:"status"`
}

type OrderResponse struct {
	ID         uint           `json:"id"`
	UserID     uint           `json:"user_id"`
	TotalPrice float64        `json:"total_price"`
	Status     string         `json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	Items      []OrderItemDTO `json:"items"`
}
