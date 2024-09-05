package user

import "bitroom/types"

type UserStoreInterface interface {
	EditUserData(data *EditUser, userId uint) *types.CustomError
	CheckUserCompletedData(id uint) *types.CustomError
}

type UserServiceInterface interface {
	EditUserData(data *EditUser, userId uint) *types.CustomError
}
