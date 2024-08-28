package news_model

import "gorm.io/gorm"

type News struct {
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Status      string `gorm:"default:'InProgress'"`
	gorm.Model
}
