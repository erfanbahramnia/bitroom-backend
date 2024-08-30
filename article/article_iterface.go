package article

import (
	article_model "bitroom/models/article"
	"bitroom/types"
)

type ArticleStoreInterface interface {
	AddArticle(data *NewArticle) (*Article, *types.CustomError)
	GetArticles() ([]MinimumArticle, *types.CustomError)
	GetArticleById(id uint) (*article_model.Article, *types.CustomError)
	GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError)
	EditArticle() (*Article, *types.CustomError)
	DeleteArticle(id uint) *types.CustomError
	GetPopularArticles() ([]MinimumArticle, *types.CustomError)
	LikeArticle(userId, articleId uint) *types.CustomError
	DislikeArticle(userId, articleId uint) *types.CustomError
	AddCommentToArticle(data *UserComment) *types.CustomError
	EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError)
	DeleteArticleComment(userId, commentId uint) *types.CustomError
	CheckArticleExist(id uint) (bool, *types.CustomError)
	CheckCategoryExist(id uint) (bool, *types.CustomError)
}

type ArticleServiceInterface interface {
	AddArticle(data *NewArticle) (*Article, *types.CustomError)
	GetArticles() ([]MinimumArticle, *types.CustomError)
	GetArticleById(id uint) (*article_model.Article, *types.CustomError)
	GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError)
	EditArticle() (*Article, *types.CustomError)
	DeleteArticle(id uint) *types.CustomError
	GetPopularArticles() ([]MinimumArticle, *types.CustomError)
	LikeArticle(userId, articleId uint) *types.CustomError
	DislikeArticle(userId, articleId uint) *types.CustomError
	AddCommentToArticle(data *UserComment) *types.CustomError
	EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError)
	DeleteArticleComment(userId, commentId uint) *types.CustomError
}
