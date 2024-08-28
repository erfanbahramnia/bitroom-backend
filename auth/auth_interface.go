package auth

import (
	user_model "bitroom/models/user"
	"bitroom/types"
)

type AuthStoreInterface interface {
	CheckUserExist(phone string) (bool, *types.CustomError)
	CreateNewUser(phone string) (*user_model.User, *types.CustomError)
	GetUserByPhone(phone string) (*user_model.User, *types.CustomError)
}

type AuthServiceInterface interface {
	Login(user LoginCredential) (*user_model.User, *types.CustomError)
	OtpGeneratingForRegister(phone string) (string, *types.CustomError)
	ValidateOtpRegister(data ValidateOtpRegistering) (*user_model.User, *types.CustomError)
}
