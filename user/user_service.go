package user

import (
	"bitroom/types"
	"sync"
)

type UserService struct {
	store UserStoreInterface
}

func NewUserSerivce(store UserStoreInterface) *UserService {
	return &UserService{
		store: store,
	}
}

// -----------------------------------------------------------------------------------------------------------

func (u *UserService) EditUserData(data *EditUser, userId uint) *types.CustomError {
	var wg sync.WaitGroup
	// update
	errChan := make(chan *types.CustomError, 5)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := u.store.EditUserData(data, userId)
		errChan <- err
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	if err := <-errChan; err != nil {
		return err
	}

	return nil
}

// -----------------------------------------------------------------------------------------------------------

func (u *UserService) ChangePaasword(phone, password string) *types.CustomError {
	// update pass
	err := u.store.ChangePaasword(phone, password)
	return err
}
