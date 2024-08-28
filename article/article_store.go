package article

import (
	article_model "bitroom/models/article"
	"bitroom/types"

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

func (a *ArticleStore) AddArticle(data *NewArticle) (*article_model.Article, *types.CustomError) {
	return nil, nil
}

func (a *ArticleStore) GetArticles() ([]article_model.Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) GetArticleById(id uint) (*article_model.Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) GetArticlesByCategory(categoryId uint) ([]article_model.Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) EditArticle() (*article_model.Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) DeleteArticle(id uint) *types.CustomError {

	return nil
}

func (a *ArticleStore) GetPopularArticles() ([]article_model.Article, *types.CustomError) {

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

func (a *ArticleStore) EditCommentToArticle(data *UserComment, commentId uint) (*article_model.ArticleComment, *types.CustomError) {

	return nil, nil
}

func (a *ArticleStore) DeleteCommentToArticle(userId, commentId uint) *types.CustomError {

	return nil
}
