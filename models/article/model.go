package article_model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string           `gorm:"not null"`
	Description string           `gorm:"not null"`
	Summary     string           `gorm:"not null"`
	Image       string           `gorm:"not null"`
	Status      string           `gorm:"default:'InProgress'"`
	Comments    []ArticleComment `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ArticleComment struct {
	gorm.Model
	Comment   string `gorm:"not null"`
	ArticleID uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
}
