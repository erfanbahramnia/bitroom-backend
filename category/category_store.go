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

// --------------------------------------------------------------------------------------------------------

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

// --------------------------------------------------------------------------------------------------------

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

// --------------------------------------------------------------------------------------------------------

func (c *CategoryStore) EditCategory(data EditCategory) *types.CustomError {
	// check category exist
	exists, err := c.CheckCategoryExist(data.ID)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// update
	var category category_model.Category
	c.db.First(&category, data.ID)
	category.Name = data.Name
	c.db.Save(&category)
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------

func (c *CategoryStore) DeleteCategory(id uint) *types.CustomError {
	// check category exist
	exists, err := c.CheckCategoryExist(id)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// get all children
	var categories []category_model.Category
	query := `
		WITH RECURSIVE category_tree AS (
			SELECT * FROM categories WHERE id = ?
			UNION ALL
			SELECT c.* FROM categories c
			INNER JOIN category_tree ct ON ct.id = c.parent_id
		)
		SELECT * FROM category_tree;
	`
	if err := c.db.Raw(query, id).Scan(&categories).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// get ids
	var ids []uint
	for _, category := range categories {
		ids = append(ids, category.ID)
	}
	// delete categories
	if err := c.db.Delete(&category_model.Category{}, ids).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------

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

func (c *CategoryStore) GetCategories() ([]*CategoryData, *types.CustomError) {
	// get all categories
	var categories []category_model.Category
	if err := c.db.Select("name, id, parent_id").Find(&categories).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	var data []*CategoryData
	for _, category := range categories {
		item := &CategoryData{
			ID:       category.ID,
			Name:     category.Name,
			ParentID: category.ParentID,
		}
		data = append(data, item)
	}
	// success
	return data, nil
}

// --------------------------------------------------------------------------------------------------------

func (c *CategoryStore) GetCategoriesTree() ([]*CategoryData, *types.CustomError) {
	// get all categories
	var categories []category_model.Category
	if err := c.db.Select("name, id, parent_id").Find(&categories).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// convert array to tree
	data := arrayToTree(categories)
	// success
	return data, nil
}

// --------------------------------------------------------------------------------------------------------

func (c *CategoryStore) CheckCategoryExist(id uint) (bool, *types.CustomError) {
	var exists bool
	err := c.db.Model(&category_model.Category{}).Select("count(*) > 0").Where("ID = ?", id).Scan(&exists).Error
	if err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return exists, nil
}

// --------------------------------------------------------------------------------------------------------

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
