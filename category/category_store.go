package category

import (
	"bitroom/constants"
	category_model "bitroom/models/category"
	"bitroom/types"
	"bitroom/utils"
	"net/http"

	"gorm.io/gorm"
)

type CategoryStore struct {
	db *gorm.DB
}

func NewCategoryStore(db *gorm.DB) *CategoryStore {
	return &CategoryStore{
		db: db,
	}
}

func (c *CategoryStore) AddCategory(name string) *types.CustomError {
	// create new category
	category := category_model.Category{
		Name:     name,
		ParentID: nil,
	}
	result := c.db.Create(&category)
	// check reusult
	if result.Error != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return nil
}

func (c *CategoryStore) AddChildCategory(data NewCategory) *types.CustomError {
	// check parent exist
	exists, err := c.CheckCategoryExist(*data.ParentID)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// create new category
	category := category_model.Category{
		Name:     data.Name,
		ParentID: data.ParentID,
	}
	result := c.db.Create(&category)
	// check reusult
	if result.Error != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return nil
}

func (c *CategoryStore) EditCategory(data EditCategory) *types.CustomError {
	return nil
}

func (c *CategoryStore) DeleteCategory(id uint) *types.CustomError {

	return nil
}

func (c *CategoryStore) GetCategoryById(id uint) (*category_model.Category, *types.CustomError) {

	return nil, nil
}

func (c *CategoryStore) GetCategories() ([]category_model.Category, *types.CustomError) {

	return nil, nil
}

func (c *CategoryStore) GetCategoriesTree() ([]category_model.Category, *types.CustomError) {

	return nil, nil
}

func (c *CategoryStore) CheckCategoryExist(id uint) (bool, *types.CustomError) {
	var exists bool
	err := c.db.Model(&category_model.Category{}).Select("count(*) > 0").Where("ID = ?", id).Scan(&exists).Error
	if err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return exists, nil
}
