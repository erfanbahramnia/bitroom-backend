package user

import (
	"bitroom/constants"
	user_model "bitroom/models/user"
	"bitroom/types"
	"bitroom/utils"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (u *UserStore) EditUserData(data *EditUser, userId uint) *types.CustomError {
	// get user
	var user User
	if err := u.db.Model(&user_model.User{}).First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewError(constants.NotFound, http.StatusNotFound)
		}
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// update
	if data.FirstName != nil {
		user.FirstName = *data.FirstName
	}
	if data.LastName != nil {
		user.LastName = *data.LastName
	}
	// save
	u.db.Save(&user)
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (u *UserStore) CheckUserCompletedData(id uint) *types.CustomError {
	// get user data
	var user User
	if err := u.db.Model(&user_model.User{}).First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewError(constants.NotFound, http.StatusNotFound)
		}
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// check data
	if user.FirstName == "" || user.LastName == "" {
		return utils.NewError("please complete your name", http.StatusBadRequest)
	}
	return nil
}
