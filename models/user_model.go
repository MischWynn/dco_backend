package models

import "gorm.io/gorm"

// User represents a user in the system
// @Description User data model
type User struct {
	gorm.Model
	ID       uint    `gorm:"primaryKey"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Name     string  `json:"name"`
	Phone    string  `json:"phone"`
	Address  string  `json:"address"`
	Role     string  `json:"role"`
	Orders   []Order `json:"orders"`
}
