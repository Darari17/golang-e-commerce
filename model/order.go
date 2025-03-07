package model

import "time"

type Order struct {
	ID         uint    `gorm:"primaryKey;autoIncrement"`
	UserID     uint    `gorm:"not null"`
	TotalPrice float64 `gorm:"not null;default:0"`
	Status     string  `gorm:"type:order_status;default:'pending'"`
	CreatedAt  time.Time
	User       User `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
