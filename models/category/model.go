package category_model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name      string `gorm:"not null"`
	ParentID  *uint
	Parent    *Category  `gorm:"foreignKey:ParentID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Children  []Category `gorm:"foreignKey:ParentID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	ArticleID *uint
}
