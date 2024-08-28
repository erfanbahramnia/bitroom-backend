package category

import (
	category_model "bitroom/models/category"
	"bitroom/types"
)

type CategoryStoreInterface interface {
	AddCategory(name string) *types.CustomError
	AddChildCategory(data NewCategory) *types.CustomError
	EditCategory(data EditCategory) *types.CustomError
	DeleteCategory(id uint) *types.CustomError
	GetCategoryById(id uint) (*category_model.Category, *types.CustomError)
	GetCategories() ([]category_model.Category, *types.CustomError)
	GetCategoriesTree() ([]category_model.Category, *types.CustomError)
	CheckCategoryExist(id uint) (bool, *types.CustomError)
}

type CategoryServiceInterface interface {
	AddCategory(category NewCategory) *types.CustomError
	EditCategory(data EditCategory) *types.CustomError
	DeleteCategory(id uint) *types.CustomError
	GetCategoryById(id uint) (*category_model.Category, *types.CustomError)
	GetCategories() ([]category_model.Category, *types.CustomError)
	GetCategoriesTree() ([]category_model.Category, *types.CustomError)
}
