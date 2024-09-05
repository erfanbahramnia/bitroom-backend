package article

import (
	article_model "bitroom/models/article"
	category_model "bitroom/models/category"
	"bitroom/types"
)

type ArticleStoreInterface interface {
	AddArticle(data *NewArticle) (*Article, *types.CustomError)
	AddArticleProperty(data *ArticleProperty) *types.CustomError
	AddCommentToArticle(data *UserComment) *types.CustomError
	GetCategory(id uint) (*category_model.Category, *types.CustomError)
	GetArticles() ([]MinimumArticle, *types.CustomError)
	GetArticleById(id uint) (*article_model.Article, *types.CustomError)
	GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError)
	GetPopularArticles() ([]MinimumArticle, *types.CustomError)
	EditArticle(*EditArticle) *types.CustomError
	EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError)
	EditArticleProperty(data *EditArticleProperty) *types.CustomError
	DeleteArticle(id uint) *types.CustomError
	DeleteArticleProperty(id uint) *types.CustomError
	DeleteArticleComment(userId, commentId uint) *types.CustomError
	CheckArticleExist(id uint) (bool, *types.CustomError)
	CheckCategoryExist(id uint) (bool, *types.CustomError)
	CheckPropertyExists(id uint) (bool, *types.CustomError)
	LikeArticle(data *LikeOrDislikeArticle) *types.CustomError
	DislikeArticle(data *LikeOrDislikeArticle) *types.CustomError
	CheckUserDisliked(data *LikeOrDislikeArticle) (bool, *types.CustomError)
	RemoveFromDislike(data *LikeOrDislikeArticle) *types.CustomError
	CheckUserLiked(data *LikeOrDislikeArticle) (bool, *types.CustomError)
	RemoveFromLike(data *LikeOrDislikeArticle) *types.CustomError
}

type ArticleServiceInterface interface {
	AddArticle(data *NewArticle) (*Article, *types.CustomError)
	GetArticles() ([]MinimumArticle, *types.CustomError)
	GetArticleById(id uint) (*article_model.Article, *types.CustomError)
	GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError)
	EditArticle(*EditArticle) *types.CustomError
	DeleteArticle(id uint) *types.CustomError
	GetPopularArticles() ([]MinimumArticle, *types.CustomError)
	AddCommentToArticle(data *UserComment) *types.CustomError
	EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError)
	DeleteArticleComment(userId, commentId uint) *types.CustomError
	AddArticleProperty(data *ArticleProperty) *types.CustomError
	EditArticleProperty(data *EditArticleProperty) *types.CustomError
	DeleteArticleProperty(id uint) *types.CustomError
	LikeArticle(data *LikeOrDislikeArticle) *types.CustomError
	DislikeArticle(data *LikeOrDislikeArticle) *types.CustomError
}
