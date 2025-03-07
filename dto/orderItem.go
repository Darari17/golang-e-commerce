package dto

type OrderItemDTO struct {
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}
