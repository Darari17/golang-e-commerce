package model

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);unique;not null"`
	Password  string    `gorm:"type:text;not null"`
	Role      string    `gorm:"type:varchar(20);default:'user'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
