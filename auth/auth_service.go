package auth

import (
	"bitroom/constants"
	user_model "bitroom/models/user"
	"bitroom/types"
	"bitroom/utils"
	"net/http"

	"github.com/patrickmn/go-cache"
)

type AuthService struct {
	store AuthStoreInterface
}

func NewAuthService(store AuthStoreInterface) *AuthService {
	return &AuthService{
		store: store,
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *AuthService) LoginWithPassword(user *LoginCredential) (*user_model.User, *types.CustomError) {
	errorChan := make(chan *types.CustomError, 1)
	userChan := make(chan *user_model.User, 1)
	go func() {
		user, err := a.store.GetUserByPhone(user.Phone)
		if err != nil {
			errorChan <- err
		}
		userChan <- user
	}()

	select {
	case userData := <-userChan:
		// check user have password or not
		if len(userData.Password) < 1 {
			return nil, utils.NewError(constants.UserHasNotPassword, http.StatusBadRequest)
		}
		// check password
		isSame := utils.CheckPasswordWithHash(user.Password, userData.Password)
		if !isSame {
			return nil, utils.NewError(constants.IncorrectPassword, http.StatusBadRequest)
		}
		return userData, nil
	case err := <-errorChan:
		return nil, err
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *AuthService) OtpGeneratingForRegister(phone string) (string, *types.CustomError) {
	errChan := make(chan *types.CustomError, 1)
	existsChan := make(chan bool, 1)

	go func() {
		exists, err := a.store.CheckUserExist(phone)
		if err != nil {
			errChan <- err
		}
		existsChan <- exists
	}()

	select {
	case err := <-errChan:
		return "", err
	case isExist := <-existsChan:
		if isExist {
			return "", utils.NewError(constants.UserAlreadyExist, http.StatusBadRequest)
		}
	}

	// saving otp with user phone in cache
	otp := utils.GenerateOtp()
	c := utils.GetCache()
	// delete previous otp
	c.Delete(phone)
	// save new otp
	c.Set(phone, otp, cache.DefaultExpiration)

	return otp, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *AuthService) OtpGeneratingForLogin(phone string) (string, *types.CustomError) {
	errChan := make(chan *types.CustomError, 1)
	existsChan := make(chan bool, 1)

	go func() {
		exists, err := a.store.CheckUserExist(phone)
		if err != nil {
			errChan <- err
		}
		existsChan <- exists
	}()

	select {
	case err := <-errChan:
		return "", err
	case isExist := <-existsChan:
		if !isExist {
			return "", utils.NewError(constants.NotFound, http.StatusNotFound)
		}
	}

	// saving otp with user phone in cache
	otp := utils.GenerateOtp()
	c := utils.GetCache()
	// delete previous otp
	c.Delete(phone)
	// save new otp
	c.Set(phone, otp, cache.DefaultExpiration)
	return otp, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *AuthService) ValidateOtpLogin(data ValidateOtp) (*user_model.User, *types.CustomError) {
	// get valid otp
	c := utils.GetCache()
	validOtp, found := c.Get(data.Phone)
	if !found {
		return nil, utils.NewError(constants.OtpExpired, http.StatusBadRequest)
	}
	// check otp is valid
	if validOtp != data.Otp {
		return nil, utils.NewError(constants.OtpInvalid, http.StatusBadRequest)
	}

	// create new user
	errChan := make(chan *types.CustomError, 1)
	userIdChan := make(chan *user_model.User, 1)
	go func() {
		user, err := a.store.GetUserByPhone(data.Phone)
		if err != nil {
			errChan <- err
		}
		userIdChan <- user
	}()

	select {
	case err := <-errChan:
		return nil, err
	case user := <-userIdChan:
		return user, nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *AuthService) ValidateOtpRegister(data ValidateOtp) (*user_model.User, *types.CustomError) {
	// get valid otp
	c := utils.GetCache()
	validOtp, found := c.Get(data.Phone)
	if !found {
		return nil, utils.NewError(constants.OtpExpired, http.StatusBadRequest)
	}
	// check otp is valid
	if validOtp != data.Otp {
		return nil, utils.NewError(constants.OtpInvalid, http.StatusBadRequest)
	}

	// create new user
	errChan := make(chan *types.CustomError, 1)
	userIdChan := make(chan *user_model.User, 1)
	go func() {
		user, err := a.store.CreateNewUser(data.Phone)
		if err != nil {
			errChan <- err
		}
		userIdChan <- user
	}()

	select {
	case err := <-errChan:
		return nil, err
	case user := <-userIdChan:
		return user, nil
	}
}
