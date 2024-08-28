package course_model

import "gorm.io/gorm"

type Course struct {
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      string `gorm:"default:'InProgress'"`
	gorm.Model
}
