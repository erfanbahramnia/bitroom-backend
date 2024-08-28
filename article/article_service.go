package article

import (
	article_model "bitroom/models/article"
	"bitroom/types"
)

type ArticleService struct {
	store ArticleStoreInterface
}

func GetNewArticeService(store ArticleStoreInterface) *ArticleService {
	return &ArticleService{
		store: store,
	}
}

func (a *ArticleService) AddArticle(data *NewArticle) (*article_model.Article, *types.CustomError) {
	return nil, nil
}

func (a *ArticleService) GetArticles() ([]article_model.Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) GetArticleById(id uint) (*article_model.Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) GetArticlesByCategory(categoryId uint) ([]article_model.Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) EditArticle() (*article_model.Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) DeleteArticle(id uint) *types.CustomError {

	return nil
}

func (a *ArticleService) GetPopularArticles() ([]article_model.Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) LikeArticle(userId, articleId uint) *types.CustomError {

	return nil
}

func (a *ArticleService) DislikeArticle(userId, articleId uint) *types.CustomError {

	return nil
}

func (a *ArticleService) AddCommentToArticle(data *UserComment) *types.CustomError {

	return nil
}

func (a *ArticleService) EditCommentToArticle(data *UserComment, commentId uint) (*article_model.ArticleComment, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) DeleteCommentToArticle(userId, commentId uint) *types.CustomError {

	return nil
}
