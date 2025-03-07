package dto

import "time"

type ProductRequest struct {
	Name     string  `json:"name" validate:"required,max=255"`
	Price    float64 `json:"price" validate:"required"`
	Stock    uint    `json:"stock" validate:"required"`
	Category string  `json:"category" validate:"required,max=255"`
}

type ProductUpdateRequest struct {
	Name     string   `json:"name" validate:"omitempty,max=255"`
	Price    *float64 `json:"price" validate:"omitempty"`
	Stock    *uint    `json:"stock" validate:"omitempty"`
	Category string   `json:"category" validate:"omitempty,max=255"`
}

type ProductResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     uint      `json:"stock"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
