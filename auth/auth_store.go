package auth

import (
	user_model "bitroom/models/user"
	"bitroom/types"
	"bitroom/utils"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type AuthStore struct {
	db *gorm.DB
}

func NewAuthStore(db *gorm.DB) *AuthStore {
	return &AuthStore{
		db: db,
	}
}

func (a *AuthStore) CheckUserExist(phone string) (bool, *types.CustomError) {
	var exists bool
	err := a.db.Model(&user_model.User{}).Select("count(*) > 0").Where("phone = ?", phone).Scan(&exists).Error
	if err != nil {
		fmt.Println(err)
		return false, utils.NewError("internal server error", http.StatusInternalServerError)
	}
	return exists, nil
}

func (a *AuthStore) CreateNewUser(phone string) (*user_model.User, *types.CustomError) {
	// save new user
	newUser := user_model.User{
		Phone: phone,
	}
	if err := a.db.Create(&newUser).Error; err != nil {
		fmt.Println(err)
		return nil, utils.NewError("could not save user", http.StatusInternalServerError)
	}

	// return new user id
	return &newUser, nil
}

func (a *AuthStore) GetUserByPhone(phone string) (*user_model.User, *types.CustomError) {
	var user user_model.User

	if err := a.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, utils.NewError("user not found", http.StatusNotFound)
		}
		return nil, utils.NewError("internal server error", http.StatusInternalServerError)
	}

	return &user, nil
}
