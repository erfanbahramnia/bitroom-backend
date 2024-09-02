package utils

import (
	"bitroom/constants"
	"bitroom/types"
	"math/rand"
	"sync"

	"golang.org/x/crypto/bcrypt"
)

func NewError(msg string, code int) *types.CustomError {
	return &types.CustomError{
		Message: msg,
		Code:    code,
	}
}

func GenerateOtp() string {
	result := make([]byte, constants.OtpLength)
	for i := range result {
		result[i] = constants.Numbers[rand.Intn(len(constants.Numbers))]
	}
	return string(result)
}

func GenerateRandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = constants.Letters[rand.Intn(len(constants.Letters))]
	}
	return string(result)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordWithHash(passowrd, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passowrd))
	return err == nil
}

func CheckExistence(id uint, checker types.ExsitenceChecker, bufferSize int) (bool, *types.CustomError) {
	var wg sync.WaitGroup

	errChan := make(chan *types.CustomError, bufferSize)
	existsChan := make(chan bool, bufferSize)
	wg.Add(1)
	go func() {
		defer wg.Done()
		exists, err := checker(id)
		if err != nil {
			errChan <- err
			return
		}
		existsChan <- exists
	}()

	go func() {
		wg.Wait()
		close(errChan)
		close(existsChan)
	}()

	select {
	case err := <-errChan:
		return false, err
	case exists := <-existsChan:
		return exists, nil
	}
}
