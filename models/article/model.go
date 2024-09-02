package article_model

import (
	category_model "bitroom/models/category"
)

type Article struct {
	ID          uint                    `gorm:"primaryKey"`
	Title       string                  `gorm:"not null"`
	Description string                  `gorm:"not null"`
	Summary     string                  `gorm:"not null"`
	Image       string                  `gorm:"not null"`
	Status      string                  `gorm:"default:'InProgress'"`
	Properties  []ArticleProperty       `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Comments    []ArticleComment        `gorm:"foreignKey:ArticleID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category    category_model.Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CategoryID  uint                    `gorm:"not null"`
}

type ArticleComment struct {
	ID        uint   `gorm:"primaryKey"`
	Comment   string `gorm:"type:text;not null"`
	ArticleID uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
}

type ArticleProperty struct {
	ID          uint   `gorm:"primaryKey"`
	Description string `gorm:"type:text;not null;column:description"`
	Image       string `gorm:"type:varchar(255);default:''"`
	ArticleID   uint   `gorm:"not null"`
}
