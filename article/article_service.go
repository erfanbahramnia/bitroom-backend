package article

import (
	"bitroom/constants"
	article_model "bitroom/models/article"
	"bitroom/types"
	"bitroom/utils"
	"net/http"
	"sync"
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
	var wg sync.WaitGroup
	// check category exist
	exists, err := utils.CheckExistence(data.Category, a.store.CheckCategoryExist, 1)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// add new article
	articleChan := make(chan *Article, 1)
	addErrChan := make(chan *types.CustomError, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		article, err := a.store.AddArticle(data)
		if err != nil {
			addErrChan <- err
			return
		}
		articleChan <- article
	}()

	go func() {
		wg.Wait()
		close(articleChan)
		close(addErrChan)
	}()

	select {
	case article := <-articleChan:
		return article, nil
	case err := <-addErrChan:
		return nil, err
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetArticles() ([]MinimumArticle, *types.CustomError) {
	articlesChan := make(chan []MinimumArticle, 20)
	errChan := make(chan *types.CustomError, 20)
	go func() {
		articles, err := a.store.GetArticles()
		if err != nil {
			errChan <- err
			return
		}
		articlesChan <- articles
	}()
	select {
	case err := <-errChan:
		return nil, err
	case articles := <-articlesChan:
		return articles, nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetArticleById(id uint) (*article_model.Article, *types.CustomError) {
	var wg sync.WaitGroup
	// check article exist
	exists, err := utils.CheckExistence(id, a.store.CheckArticleExist, 20)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// get article
	articleChan := make(chan *article_model.Article, 20)
	errChan := make(chan *types.CustomError, 20)
	wg.Add(1)
	go func() {
		defer wg.Done()
		article, err := a.store.GetArticleById(id)
		if err != nil {
			errChan <- err
			return
		}
		articleChan <- article
	}()

	go func() {
		wg.Wait()
		close(articleChan)
		close(errChan)
	}()

	select {
	case err := <-errChan:
		return nil, err
	case article := <-articleChan:
		return article, nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError) {
	var wg sync.WaitGroup
	// check category exist
	exists, err := utils.CheckExistence(categoryId, a.store.CheckCategoryExist, 20)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// get articles
	articlesChan := make(chan []MinimumArticle, 20)
	errChan := make(chan *types.CustomError, 20)
	wg.Add(1)
	go func() {
		defer wg.Done()
		artilces, err := a.store.GetArticlesByCategory(categoryId)
		if err != nil {
			errChan <- err
			return
		}
		articlesChan <- artilces
	}()

	go func() {
		wg.Wait()
		close(articlesChan)
		close(errChan)
	}()

	select {
	case err := <-errChan:
		return nil, err
	case articles := <-articlesChan:
		return articles, nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) EditArticle(article *EditArticle) *types.CustomError {
	var wg sync.WaitGroup

	// check article exists
	exists, err := utils.CheckExistence(*article.Id, a.store.CheckArticleExist, 1)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// check category exists
	if article.Category != nil {
		exists, err := utils.CheckExistence(*article.Category, a.store.CheckCategoryExist, 1)
		if err != nil {
			return err
		}
		if !exists {
			return utils.NewError(constants.NotFound, http.StatusNotFound)
		}
	}

	// update
	editErrChan := make(chan *types.CustomError, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.store.EditArticle(article)
		editErrChan <- err
	}()

	go func() {
		wg.Wait()
		close(editErrChan)
	}()

	// success
	select {
	case err := <-editErrChan:
		return err
	default:
		return nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DeleteArticle(id uint) *types.CustomError {
	var wg sync.WaitGroup
	// check article exists
	exists, err := utils.CheckExistence(id, a.store.CheckArticleExist, 1)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}
	// delete images of article

	// delete article
	errChan := make(chan *types.CustomError, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.store.DeleteArticle(id)
		errChan <- err
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) AddArticleProperty(data *ArticleProperty) *types.CustomError {
	var wg sync.WaitGroup
	// check article exists
	exists, err := utils.CheckExistence(data.ArticleID, a.store.CheckArticleExist, 1)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// add property
	errChan := make(chan *types.CustomError, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.store.AddArticleProperty(data)
		errChan <- err
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) EditArticleProperty(data *EditArticleProperty) *types.CustomError {
	var wg sync.WaitGroup
	// check property exists
	exists, err := utils.CheckExistence(data.PropertyID, a.store.CheckPropertyExists, 1)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// edit
	errChan := make(chan *types.CustomError, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.store.EditArticleProperty(data)
		errChan <- err
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DeleteArticleProperty(id uint) *types.CustomError {
	var wg sync.WaitGroup
	// check property exists
	exists, err := utils.CheckExistence(id, a.store.CheckPropertyExists, 1)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// delete
	errChan := make(chan *types.CustomError, 1)
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := a.store.DeleteArticleProperty(id)
		errChan <- err
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
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
