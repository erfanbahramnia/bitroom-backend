package auth

import (
	"bitroom/models"
	"bitroom/types"
)

type AuthStoreInterface interface {
	CheckUserExist(phone string) (bool, *types.CustomError)
	CreateNewUser(phone string) (*models.User, *types.CustomError)
	GetUserByPhone(phone string) (*models.User, *types.CustomError)
}

type AuthServiceInterface interface {
	Login(user LoginCredential) (*models.User, *types.CustomError)
	OtpGeneratingForRegister(phone string) (string, *types.CustomError)
	ValidateOtpRegister(data ValidateOtpRegistering) (*models.User, *types.CustomError)
}
