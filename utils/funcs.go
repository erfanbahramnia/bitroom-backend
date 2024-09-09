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

func ReactionChecker(data *types.LikeOrDislikeArticle, checker types.ReactionChecker) (bool, *types.CustomError) {
	var wg sync.WaitGroup

	reactionChan := make(chan bool, 1)
	errChan := make(chan *types.CustomError, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		reaction, err := checker(data)
		if err != nil {
			errChan <- err
			return
		}
		reactionChan <- reaction
	}()

	go func() {
		wg.Wait()
		close(reactionChan)
		close(errChan)
	}()

	select {
	case err := <-errChan:
		return false, err
	case reactoin := <-reactionChan:
		return reactoin, nil
	}
}

func ReactionUpdator(data *types.LikeOrDislikeArticle, remover types.ReactionRemover) *types.CustomError {
	var wg sync.WaitGroup

	errChan := make(chan *types.CustomError, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := remover(data)
		errChan <- err
	}()

	err := <-errChan
	return err
}
func MapStringInterface(field string, input any) map[string]interface{} {
	return map[string]interface{}{
		field: input,
	}
}
