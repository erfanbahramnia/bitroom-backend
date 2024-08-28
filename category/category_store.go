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

func (c *CategoryStore) GetCategoryById(id uint) ([]*CategoryData, *types.CustomError) {
	var categories []category_model.Category
	// recursive query for getting all children of selected category
	recursiveQuery := `
		WITH RECURSIVE category_tree as (
			SELECT * FROM categories WHERE id = ?
			UNION ALL
			SELECT c.* FROM categories c
			INNER JOIN category_tree ct ON ct.id = c.Parent_id
		)
		SELECT * FROM category_tree;
	`
	if err := c.db.Raw(recursiveQuery, id).Scan(&categories).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// convert array to tree
	data := arrayToTree(categories)
	// success
	return data, nil
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

func arrayToTree(categories []category_model.Category) []*CategoryData {
	categoryMap := make(map[uint]*CategoryData)
	var rootCategories []*CategoryData

	for _, category := range categories {
		categoryMap[category.ID] = &CategoryData{
			Name:     category.Name,
			ParentID: category.ParentID,
			ID:       category.ID,
			Children: []*CategoryData{},
		}
	}

	for _, category := range categories {
		node := categoryMap[category.ID]
		if node.ParentID != nil {
			parentNode := categoryMap[*node.ParentID]
			parentNode.Children = append(parentNode.Children, node)
		} else {
			rootCategories = append(rootCategories, node)
		}
	}

	return rootCategories
}
