package article

import (
	"bitroom/constants"
	article_model "bitroom/models/article"
	"bitroom/types"
	"bitroom/utils"
	"net/http"
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
	// check category exists
	exists, existsErr := a.store.CheckCategoryExist(data.Category)
	if existsErr != nil {
		return nil, existsErr
	}
	if !exists {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	article, err := a.store.AddArticle(data)
	if err != nil {
		return nil, err
	}
	return article, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetArticles() ([]MinimumArticle, *types.CustomError) {
	articles, err := a.store.GetArticles()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetArticleById(id uint) (*article_model.Article, *types.CustomError) {
	// check article exist
	exists, checkingErr := a.store.CheckArticleExist(id)
	if checkingErr != nil {
		return nil, checkingErr
	}
	if !exists {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// get article
	article, err := a.store.GetArticleById(id)
	if err != nil {
		return nil, err
	}
	// success
	return article, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError) {
	// check category exist
	exist, existsErr := a.store.CheckCategoryExist(categoryId)
	if existsErr != nil {
		return nil, existsErr
	}
	if !exist {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// get articles
	artilces, err := a.store.GetArticlesByCategory(categoryId)
	if err != nil {
		return nil, err
	}
	return artilces, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) EditArticle(article *EditArticle) *types.CustomError {
	// check article exists
	exists, err := a.store.CheckArticleExist(*article.Id)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// check category exists
	if article.Category != nil {
		exists, existsErr := a.store.CheckCategoryExist(*article.Category)
		if existsErr != nil {
			return existsErr
		}
		if !exists {
			return utils.NewError(constants.NotFound, http.StatusNotFound)
		}
	}
	// update
	err = a.store.EditArticle(article)
	// success
	return err
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DeleteArticle(id uint) *types.CustomError {
	// check article exist
	exists, checkingErr := a.store.CheckArticleExist(id)
	if checkingErr != nil {
		return checkingErr
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// delete images of article

	// delete article
	err := a.store.DeleteArticle(id)
	if err != nil {
		return err
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) AddArticleProperty(data *ArticleProperty) *types.CustomError {
	// check article exists
	exists, err := a.store.CheckArticleExist(data.ArticleID)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError("article not found", http.StatusNotFound)
	}
	// add property
	err = a.store.AddArticleProperty(data)
	return err
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) EditArticleProperty(data *ArticleProperty) *types.CustomError {

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DeleteArticleProperty(id uint) *types.CustomError {
	// check property exists
	exists, err := a.store.CheckPropertyExists(id)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// delete
	err = a.store.DeleteArticleProperty(id)
	return err
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetPopularArticles() ([]MinimumArticle, *types.CustomError) {

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
