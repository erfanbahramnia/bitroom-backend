package category

import (
	"bitroom/types"
)

type CategoryStoreInterface interface {
	AddCategory(name string) *types.CustomError
	AddChildCategory(data NewCategory) *types.CustomError
	EditCategory(data EditCategory) *types.CustomError
	DeleteCategory(id uint) *types.CustomError
	GetCategoryById(id uint) ([]*CategoryData, *types.CustomError)
	GetCategories() ([]*CategoryData, *types.CustomError)
	GetCategoriesTree() ([]*CategoryData, *types.CustomError)
	CheckCategoryExist(id uint) (bool, *types.CustomError)
}

type CategoryServiceInterface interface {
	AddCategory(category NewCategory) *types.CustomError
	EditCategory(data EditCategory) *types.CustomError
	DeleteCategory(id uint) *types.CustomError
	GetCategoryById(id uint) ([]*CategoryData, *types.CustomError)
	GetCategories() ([]*CategoryData, *types.CustomError)
	GetCategoriesTree() ([]*CategoryData, *types.CustomError)
}
