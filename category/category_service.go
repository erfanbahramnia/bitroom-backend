package category

import (
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
	err := c.store.EditCategory(data)
	if err != nil {
		return err
	}
	return nil
}

func (c *CategoryService) DeleteCategory(id uint) *types.CustomError {

	return nil
}

func (c *CategoryService) GetCategoryById(id uint) ([]*CategoryData, *types.CustomError) {
	// get category
	errChan := make(chan *types.CustomError, 1)
	categoryChan := make(chan []*CategoryData, 1)
	go func() {
		category, err := c.store.GetCategoryById(id)
		if err != nil {
			errChan <- err
			return
		}
		categoryChan <- category
	}()

	select {
	case err := <-errChan:
		return nil, err
	case category := <-categoryChan:
		return category, nil
	}
}

func (c *CategoryService) GetCategories() ([]*CategoryData, *types.CustomError) {
	// get categories
	errChan := make(chan *types.CustomError, 1)
	categoryChan := make(chan []*CategoryData, 1)
	go func() {
		categories, err := c.store.GetCategories()
		if err != nil {
			errChan <- err
			return
		}
		categoryChan <- categories
	}()

	select {
	case err := <-errChan:
		return nil, err
	case categories := <-categoryChan:
		return categories, nil
	}
}

func (c *CategoryService) GetCategoriesTree() ([]*CategoryData, *types.CustomError) {
	// get categories
	errChan := make(chan *types.CustomError, 1)
	categoryChan := make(chan []*CategoryData, 1)
	go func() {
		categories, err := c.store.GetCategoriesTree()
		if err != nil {
			errChan <- err
			return
		}
		categoryChan <- categories
	}()

	select {
	case err := <-errChan:
		return nil, err
	case categories := <-categoryChan:
		return categories, nil
	}
}
