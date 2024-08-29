package article

import (
	"bitroom/constants"
	article_model "bitroom/models/article"
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

func (a *ArticleStore) AddArticle(data *NewArticle) (*Article, *types.CustomError) {
	// save article
	article := article_model.Article{
		Title:       data.Title,
		Description: data.Description,
		Summary:     data.Summary,
		Image:       data.Image,
	}
	err := a.db.Create(&article).Error
	if err != nil {
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

func (a *ArticleStore) GetArticles() ([]Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) GetArticleById(id uint) (*Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) GetArticlesByCategory(categoryId uint) ([]Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) EditArticle() (*Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) DeleteArticle(id uint) *types.CustomError {

	return nil
}

func (a *ArticleStore) GetPopularArticles() ([]Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) LikeArticle(userId, articleId uint) *types.CustomError {

	return nil
}

func (a *ArticleStore) DislikeArticle(userId, articleId uint) *types.CustomError {

	return nil
}

func (a *ArticleStore) AddCommentToArticle(data *UserComment) *types.CustomError {

	return nil
}

func (a *ArticleStore) EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) DeleteArticleComment(userId, commentId uint) *types.CustomError {

	return nil
}
