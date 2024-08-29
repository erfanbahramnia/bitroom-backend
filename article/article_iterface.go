package article

import (
	"bitroom/types"
)

type ArticleStoreInterface interface {
	AddArticle(data *NewArticle) (*Article, *types.CustomError)
	GetArticles() ([]Article, *types.CustomError)
	GetArticleById(id uint) (*Article, *types.CustomError)
	GetArticlesByCategory(categoryId uint) ([]Article, *types.CustomError)
	EditArticle() (*Article, *types.CustomError)
	DeleteArticle(id uint) *types.CustomError
	GetPopularArticles() ([]Article, *types.CustomError)
	LikeArticle(userId, articleId uint) *types.CustomError
	DislikeArticle(userId, articleId uint) *types.CustomError
	AddCommentToArticle(data *UserComment) *types.CustomError
	EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError)
	DeleteArticleComment(userId, commentId uint) *types.CustomError
}

type ArticleServiceInterface interface {
	AddArticle(data *NewArticle) (*Article, *types.CustomError)
	GetArticles() ([]Article, *types.CustomError)
	GetArticleById(id uint) (*Article, *types.CustomError)
	GetArticlesByCategory(categoryId uint) ([]Article, *types.CustomError)
	EditArticle() (*Article, *types.CustomError)
	DeleteArticle(id uint) *types.CustomError
	GetPopularArticles() ([]Article, *types.CustomError)
	LikeArticle(userId, articleId uint) *types.CustomError
	DislikeArticle(userId, articleId uint) *types.CustomError
	AddCommentToArticle(data *UserComment) *types.CustomError
	EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError)
	DeleteArticleComment(userId, commentId uint) *types.CustomError
}
