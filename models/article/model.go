package article_model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title       string            `gorm:"not null"`
	Description string            `gorm:"not null"`
	Summary     string            `gorm:"not null"`
	Image       string            `gorm:"not null"`
	Status      string            `gorm:"default:'InProgress'"`
	Properties  []ArticleProperty `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments    []ArticleComment  `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type ArticleComment struct {
	gorm.Model
	Comment   string `gorm:"not null"`
	ArticleID uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
}

type ArticleProperty struct {
	gorm.Model
	Text      string `gorm:"not null"`
	Image     string `gorm:"default:''"`
	ArticleID uint   `gorm:"not null"`
}
