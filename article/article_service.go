package article

import (
	article_model "bitroom/models/article"
	"bitroom/types"
)

type ArticleService struct {
	store ArticleStoreInterface
}

func NewArticleService(store ArticleStoreInterface) *ArticleService {
	return &ArticleService{
		store: store,
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) AddArticle(data *NewArticle) (*Article, *types.CustomError) {
	article, err := a.store.AddArticle(data)
	if err != nil {
		return nil, err
	}
	return article, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetArticles() ([]MinimumArtilce, *types.CustomError) {
	articles, err := a.store.GetArticles()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetArticleById(id uint) (*article_model.Article, *types.CustomError) {
	article, err := a.store.GetArticleById(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetArticlesByCategory(categoryId uint) ([]MinimumArtilce, *types.CustomError) {

	return nil, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) EditArticle() (*Article, *types.CustomError) {

	return nil, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DeleteArticle(id uint) *types.CustomError {
	err := a.store.DeleteArticle(id)
	if err != nil {
		return err
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetPopularArticles() ([]MinimumArtilce, *types.CustomError) {

	return nil, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) LikeArticle(userId, articleId uint) *types.CustomError {

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DislikeArticle(userId, articleId uint) *types.CustomError {

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) AddCommentToArticle(data *UserComment) *types.CustomError {

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError) {

	return nil, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DeleteArticleComment(userId, commentId uint) *types.CustomError {

	return nil
}
