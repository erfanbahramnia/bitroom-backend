package article

import (
	article_model "bitroom/models/article"
	category_model "bitroom/models/category"
	"bitroom/types"
)

type ArticleStoreInterface interface {
	AddArticle(data *NewArticle) (*Article, *types.CustomError)
	AddArticleProperty(data *ArticleProperty) *types.CustomError
	GetCategory(id uint) (*category_model.Category, *types.CustomError)
	GetArticles() ([]MinimumArticle, *types.CustomError)
	GetArticlesByAdmin() ([]MinimumArticle, *types.CustomError)
	GetArticleById(id uint) (*article_model.Article, *types.CustomError)
	GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError)
	GetPopularArticles() ([]MinimumArticle, *types.CustomError)
	EditArticle(*EditArticle) *types.CustomError
	EditArticleProperty(data *EditArticleProperty) *types.CustomError
	DeleteArticle(id uint) *types.CustomError
	DeleteArticleProperty(id uint) *types.CustomError
	CheckArticleExist(id uint) (bool, *types.CustomError)
	CheckArticleAllStatusExist(id uint) (bool, *types.CustomError)
	CheckCategoryExist(id uint) (bool, *types.CustomError)
	CheckPropertyExists(id uint) (bool, *types.CustomError)
	LikeArticle(data *types.LikeOrDislikeArticle) *types.CustomError
	DislikeArticle(data *types.LikeOrDislikeArticle) *types.CustomError
	CheckUserDisliked(data *types.LikeOrDislikeArticle) (bool, *types.CustomError)
	RemoveFromDislike(data *types.LikeOrDislikeArticle) *types.CustomError
	CheckUserLiked(data *types.LikeOrDislikeArticle) (bool, *types.CustomError)
	RemoveFromLike(data *types.LikeOrDislikeArticle) *types.CustomError
	AddCommentToArticle(data *NewComment) *types.CustomError
	EditArticleComment(data *EditComment) *types.CustomError
	DeleteArticleComment(data *DeleteComment) *types.CustomError
	CheckUserProvidedData(userId uint) *types.CustomError
	ChangeStatus(status string, id uint) *types.CustomError
}

type ArticleServiceInterface interface {
	AddArticle(data *NewArticle) (*Article, *types.CustomError)
	GetArticles() ([]MinimumArticle, *types.CustomError)
	GetArticlesByAdmin() ([]MinimumArticle, *types.CustomError)
	GetArticleById(id uint) (*article_model.Article, *types.CustomError)
	GetArticleByIdByAdmin(id uint) (*article_model.Article, *types.CustomError)
	GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError)
	EditArticle(*EditArticle) *types.CustomError
	DeleteArticle(id uint) *types.CustomError
	GetPopularArticles() ([]MinimumArticle, *types.CustomError)
	AddArticleProperty(data *ArticleProperty) *types.CustomError
	EditArticleProperty(data *EditArticleProperty) *types.CustomError
	DeleteArticleProperty(id uint) *types.CustomError
	LikeArticle(data *types.LikeOrDislikeArticle) *types.CustomError
	DislikeArticle(data *types.LikeOrDislikeArticle) *types.CustomError
	AddCommentToArticle(data *NewComment) *types.CustomError
	EditArticleComment(data *EditComment) *types.CustomError
	DeleteArticleComment(data *DeleteComment) *types.CustomError
	ChangeStatus(status string, id uint) *types.CustomError
}
