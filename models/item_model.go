package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	Name         string
	Desc         string
	Price        float64
	Photo        []byte
	CategoryID   uint
	Category     Category
	OrderDetails []OrderDetail `gorm:"foreignKey:ItemID"`
}
