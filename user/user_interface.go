package user

import "bitroom/types"

type UserStoreInterface interface {
	EditUserData(data *EditUser, userId uint) *types.CustomError
	CheckUserCompletedData(id uint) *types.CustomError
	ChangePaasword(phone, password string) *types.CustomError
}

type UserServiceInterface interface {
	EditUserData(data *EditUser, userId uint) *types.CustomError
	ChangePaasword(phone, password string) *types.CustomError
}
