package article

import (
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

func (a *ArticleService) AddArticle(data *NewArticle) (*Article, *types.CustomError) {
	article, err := a.store.AddArticle(data)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (a *ArticleService) GetArticles() ([]Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) GetArticleById(id uint) (*Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) GetArticlesByCategory(categoryId uint) ([]Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) EditArticle() (*Article, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) DeleteArticle(id uint) *types.CustomError {

	return nil
}

func (a *ArticleService) GetPopularArticles() ([]Article, *types.CustomError) {

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

func (a *ArticleService) EditArticleComment(data *UserComment, commentId uint) (*ArticleComment, *types.CustomError) {

	return nil, nil
}

func (a *ArticleService) DeleteArticleComment(userId, commentId uint) *types.CustomError {

	return nil
}
