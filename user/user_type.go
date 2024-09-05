package user

type EditUser struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}

type User struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `gorm:"default:''"`
	LastName  string `gorm:"default:''"`
	Role      string `gorm:"default:'user'"`
	Phone     string `gorm:"not null;unique"`
	Password  string `gorm:"default:''"`
}
