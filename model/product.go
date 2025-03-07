package model

import "time"

type Product struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	Name      string  `gorm:"type:varchar(255);not null"`
	Price     float64 `gorm:"type:decimal(10,2);not null"`
	Stock     uint    `gorm:"not null;default:0"`
	Category  string  `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
}
