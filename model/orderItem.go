package model

type OrderItem struct {
	ID        uint    `gorm:"primaryKey;autoIncrement"`
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Subtotal  float64 `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Product   Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
