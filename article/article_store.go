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

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) AddArticle(data *NewArticle) (*Article, *types.CustomError) {
	// get category
	var category category_model.Category
	if err := a.db.Find(&category, data.Category).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
		}
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// save article
	article := article_model.Article{
		Title:       data.Title,
		Description: data.Description,
		Summary:     data.Summary,
		Image:       data.Image,
		Category:    category,
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

func (a *ArticleStore) GetArticles() ([]MinimumArtilce, *types.CustomError) {
	var articles []MinimumArtilce
	if err := a.db.Model(&article_model.Article{}).Select("title, summary, image, id").Find(&articles).Error; err != nil {
		return nil, utils.NewError("could not retrive data", http.StatusInternalServerError)
	}
	// success
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetArticleById(id uint) (*article_model.Article, *types.CustomError) {
	// check article exist
	exists, checkingErr := a.CheckArticleExist(id)
	if checkingErr != nil {
		return nil, checkingErr
	}
	if !exists {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}
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

func (a *ArticleStore) GetArticlesByCategory(categoryId uint) ([]MinimumArtilce, *types.CustomError) {

	return nil, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) EditArticle() (*Article, *types.CustomError) {

	return nil, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DeleteArticle(id uint) *types.CustomError {
	// check article exist
	exists, checkingErr := a.CheckArticleExist(id)
	if checkingErr != nil {
		return checkingErr
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// delete
	if err := a.db.Delete(&article_model.Article{}, id).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetPopularArticles() ([]MinimumArtilce, *types.CustomError) {

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
