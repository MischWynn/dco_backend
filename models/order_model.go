package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID       uint
	User         User
	TotalAmount  float64
	Status       string
	Method       string
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID"`
}
