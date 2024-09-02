package category_model

type Category struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	ParentID *uint
	Parent   *Category  `gorm:"foreignKey:ParentID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	Children []Category `gorm:"foreignKey:ParentID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}
