package article

import (
	"bitroom/constants"
	article_model "bitroom/models/article"
	category_model "bitroom/models/category"
	"bitroom/types"
	"bitroom/utils"
	"net/http"

	"gorm.io/gorm"
)

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) *ArticleStore {
	return &ArticleStore{
		db: db,
	}
}

func (a *ArticleStore) GetCategory(id uint) (*category_model.Category, *types.CustomError) {
	// get category
	var category category_model.Category
	if err := a.db.Find(&category, id).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return &category, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) AddArticle(data *NewArticle) (*Article, *types.CustomError) {
	// get category
	category, err := a.GetCategory(data.Category)
	if err != nil {
		return nil, err
	}
	// save article
	article := article_model.Article{
		Title:       data.Title,
		Description: data.Description,
		Summary:     data.Summary,
		Image:       data.Image,
		Category:    *category,
	}
	if err := a.db.Create(&article).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return &Article{
		Title:       article.Title,
		Description: article.Description,
		Summary:     article.Summary,
		Image:       article.Image,
		ID:          article.ID,
		Status:      article.Status,
	}, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetArticles() ([]MinimumArticle, *types.CustomError) {
	var articles []MinimumArticle
	if err := a.db.Model(&article_model.Article{}).Select("title, summary, image, id").Find(&articles).Error; err != nil {
		return nil, utils.NewError("could not retrive data", http.StatusInternalServerError)
	}
	// success
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetArticleById(id uint) (*article_model.Article, *types.CustomError) {
	// get article
	var article article_model.Article
	err := a.db.Preload("Properties").Preload("Comments").Preload("Category").First(&article, id).Error
	if err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return &article, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError) {
	// recursive query for getting all children of selected category
	recursiveQuery := `
		WITH RECURSIVE category_tree as (
			SELECT * FROM categories WHERE id = ?
			UNION ALL
			SELECT c.* FROM categories c
			INNER JOIN category_tree ct ON ct.id = c.Parent_id
		)
		SELECT id FROM category_tree;
	`
	var ids []uint
	if err := a.db.Model(&category_model.Category{}).Raw(recursiveQuery, categoryId).Select("id").Scan(&ids).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	// get articles
	var articles []MinimumArticle
	if err := a.db.Model(&article_model.Article{}).
		Select("articles.title, articles.summary, articles.image, articles.id").
		Joins("LEFT JOIN categories ON categories.id = articles.category_id").
		Where("categories.id IN ?", ids).
		Scan(&articles).Error; err != nil {
		return nil, utils.NewError("could not retrive data", http.StatusInternalServerError)
	}
	// success
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) EditArticle(editedArticle *EditArticle) *types.CustomError {
	article := article_model.Article{}
	// get article
	if err := a.db.First(&article, *editedArticle.Id).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	// update
	if editedArticle.Title != nil {
		article.Title = *editedArticle.Title
	}
	if editedArticle.Description != nil {
		article.Description = *editedArticle.Description
	}
	if editedArticle.Summary != nil {
		article.Summary = *editedArticle.Summary
	}
	if editedArticle.Status != nil {
		article.Status = *editedArticle.Status
	}
	if editedArticle.Image != nil {
		article.Image = *editedArticle.Image
	}
	if editedArticle.Category != nil {
		// get category
		category, err := a.GetCategory(*editedArticle.Category)
		if err != nil {
			return err
		}
		article.Category = *category
	}

	// save changes
	if err := a.db.Save(&article).Error; err != nil {
		return utils.NewError("failed to update", http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DeleteArticle(id uint) *types.CustomError {
	// delete
	if err := a.db.Delete(&article_model.Article{}, id).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) AddArticleProperty(data *ArticleProperty) *types.CustomError {
	// add new property
	property := article_model.ArticleProperty{
		Image:       *data.Image,
		Description: data.Description,
		ArticleID:   data.ArticleID,
	}
	if err := a.db.Create(&property).Error; err != nil {
		return utils.NewError("could not add new property", http.StatusInternalServerError)
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) EditArticleProperty(data *EditArticleProperty) *types.CustomError {
	var property article_model.ArticleProperty
	if err := a.db.First(&property, data.PropertyID).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	// update
	if data.Description != nil {
		property.Description = *data.Description
	}
	if data.Image != nil {
		property.Image = *data.Image
	}

	// save changes
	if err := a.db.Save(property).Error; err != nil {
		return utils.NewError("failed to update", http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DeleteArticleProperty(id uint) *types.CustomError {
	if err := a.db.Delete(&article_model.ArticleProperty{}, id).Error; err != nil {
		return utils.NewError(constants.CouldNotDeleteItem, http.StatusInternalServerError)
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetPopularArticles() ([]MinimumArticle, *types.CustomError) {

	return nil, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) LikeArticle(userId, articleId uint) *types.CustomError {

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DislikeArticle(userId, articleId uint) *types.CustomError {

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) AddCommentToArticle(data *UserComment) *types.CustomError {

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError) {

	return nil, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DeleteArticleComment(userId, commentId uint) *types.CustomError {

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckArticleExist(id uint) (bool, *types.CustomError) {
	var exists bool
	err := a.db.Model(&article_model.Article{}).
		Select("count(*) > 0").
		Where("ID = ?", id).
		Scan(&exists).Error

	if err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	return exists, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckCategoryExist(id uint) (bool, *types.CustomError) {
	var exists bool
	err := a.db.Model(&category_model.Category{}).Select("count(*) > 0").Where("ID = ?", id).Scan(&exists).Error
	if err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return exists, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckPropertyExists(id uint) (bool, *types.CustomError) {
	var exists bool
	err := a.db.Model(&article_model.ArticleProperty{}).
		Select("count(*) > 0").
		Where("ID = ?", id).
		Scan(&exists).Error
	if err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return exists, nil
}
