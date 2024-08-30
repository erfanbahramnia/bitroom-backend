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

func (a *ArticleStore) GetArticles() ([]MinimumArtilce, *types.CustomError) {
	var articles []MinimumArtilce
	if err := a.db.Model(&article_model.Article{}).Select("title, summary, image, id").Find(&articles).Error; err != nil {
		return nil, utils.NewError("could not retrive data", http.StatusInternalServerError)
	}
	// success
	return articles, nil
}

func (a *ArticleStore) GetArticleById(id uint) (*article_model.Article, *types.CustomError) {
	// get article
	var article *article_model.Article
	err := a.db.Preload("Properties").Preload("Comments").Find(&article, id).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
		}
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return article, nil
}

func (a *ArticleStore) GetArticlesByCategory(categoryId uint) ([]MinimumArtilce, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) EditArticle() (*Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) DeleteArticle(id uint) *types.CustomError {

	return nil
}

func (a *ArticleStore) GetPopularArticles() ([]MinimumArtilce, *types.CustomError) {

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
