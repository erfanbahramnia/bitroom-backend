package user_model

import (
	article_model "bitroom/models/article"
)

type User struct {
	ID        uint                           `gorm:"primaryKey"`
	FirstName string                         `gorm:"default:''"`
	LastName  string                         `gorm:"default:''"`
	Role      string                         `gorm:"default:'user'"`
	Phone     string                         `gorm:"not null;unique"`
	Password  string                         `gorm:"default:''"`
	Comments  []article_model.ArticleComment `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
