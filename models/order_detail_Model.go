package models

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderID  uint  `gorm:"not null"`
	Order    Order `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ItemID   uint  `gorm:"not null"`
	Item     Item  `gorm:"foreignKey:ItemID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Quantity int
}
