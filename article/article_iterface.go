package article

import (
	article_model "bitroom/models/article"
	"bitroom/types"
)

type ArticleStoreInterface interface {
	AddArticle(data *NewArticle) (*article_model.Article, *types.CustomError)
	GetArticles() ([]article_model.Article, *types.CustomError)
	GetArticleById(id uint) (*article_model.Article, *types.CustomError)
	GetArticlesByCategory(categoryId uint) ([]article_model.Article, *types.CustomError)
	EditArticle() (*article_model.Article, *types.CustomError)
	DeleteArticle(id uint) *types.CustomError
	GetPopularArticles() ([]article_model.Article, *types.CustomError)
	LikeArticle(userId, articleId uint) *types.CustomError
	DislikeArticle(userId, articleId uint) *types.CustomError
	AddCommentToArticle(data *UserComment) *types.CustomError
	EditCommentToArticle(data *UserComment, commentId uint) (*article_model.ArticleComment, *types.CustomError)
	DeleteCommentToArticle(userId, commentId uint) *types.CustomError
}
