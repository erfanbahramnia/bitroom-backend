package category

import (
	category_model "bitroom/models/category"
	"bitroom/types"
	"sync"
)

type CategoryService struct {
	store CategoryStoreInterface
}

func NewCategoryService(store CategoryStoreInterface) *CategoryService {
	return &CategoryService{
		store: store,
	}
}

func (c *CategoryService) AddCategory(category NewCategory) *types.CustomError {
	var wg sync.WaitGroup
	// add new categoory
	errChan := make(chan *types.CustomError, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		var err *types.CustomError
		if category.ParentID == nil {
			err = c.store.AddCategory(category.Name)
		} else {
			err = c.store.AddChildCategory(category)
		}
		if err != nil {
			errChan <- err
		}
	}()
	wg.Wait()
	close(errChan)
	// get result
	if err, ok := <-errChan; ok {
		return err
	}
	return nil
}

func (c *CategoryService) EditCategory(data EditCategory) *types.CustomError {
	return nil
}

func (c *CategoryService) DeleteCategory(id uint) *types.CustomError {

	return nil
}

func (c *CategoryService) GetCategoryById(id uint) (*category_model.Category, *types.CustomError) {

	return nil, nil
}

func (c *CategoryService) GetCategories() ([]category_model.Category, *types.CustomError) {

	return nil, nil
}

func (c *CategoryService) GetCategoriesTree() ([]category_model.Category, *types.CustomError) {

	return nil, nil
}
