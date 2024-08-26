package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"not null;default:''"`
	LastName  string `gorm:"not null;default:''"`
	Role      string `gorm:"default:'user'"`
	Phone     string `gorm:"not null"`
	Password  string `gorm:"not null;default:''"`
}
