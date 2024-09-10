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
	articlesChan := make(chan []MinimumArticle, 1)
	errChan := make(chan *types.CustomError, 1)
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

func (a *ArticleService) GetArticlesByAdmin() ([]MinimumArticle, *types.CustomError) {
	articlesChan := make(chan []MinimumArticle, 1)
	errChan := make(chan *types.CustomError, 1)
	go func() {
		articles, err := a.store.GetArticlesByAdmin()
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
	exists, err := utils.CheckExistence(id, a.store.CheckArticleExist, 1)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// get article
	articleChan := make(chan *article_model.Article, 1)
	errChan := make(chan *types.CustomError, 1)
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

func (a *ArticleService) GetArticleByIdByAdmin(id uint) (*article_model.Article, *types.CustomError) {
	var wg sync.WaitGroup
	// check article exist
	exists, err := utils.CheckExistence(id, a.store.CheckArticleAllStatusExist, 1)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// get article
	articleChan := make(chan *article_model.Article, 1)
	errChan := make(chan *types.CustomError, 1)
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
	exists, err := utils.CheckExistence(categoryId, a.store.CheckCategoryExist, 1)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// get articles
	articlesChan := make(chan []MinimumArticle, 1)
	errChan := make(chan *types.CustomError, 1)
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
	articles, err := a.store.GetPopularArticles()
	if err != nil {
		return nil, err
	}
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) LikeArticle(data *types.LikeOrDislikeArticle) *types.CustomError {

	// check user already liked or not
	liked, err := utils.ReactionChecker(data, a.store.CheckUserLiked)
	if err != nil {
		return err
	}
	if liked {
		return utils.NewError("user already liked", http.StatusBadRequest)
	}
	// check user already disliked or not
	disliked, err := utils.ReactionChecker(data, a.store.CheckUserDisliked)
	if err != nil {
		return err
	}
	if disliked {
		// remove from dislike
		err = utils.ReactionUpdator(data, a.store.RemoveFromDislike)
		if err != nil {
			return err
		}
	}
	// update
	err = utils.ReactionUpdator(data, a.store.LikeArticle)
	if err != nil {
		return err
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DislikeArticle(data *types.LikeOrDislikeArticle) *types.CustomError {
	// check user already disliked or not
	disliked, err := utils.ReactionChecker(data, a.store.CheckUserDisliked)
	if err != nil {
		return err
	}
	if disliked {
		return utils.NewError("user already disliked", http.StatusBadRequest)
	}
	// check user already liked or not
	liked, err := utils.ReactionChecker(data, a.store.CheckUserLiked)
	if err != nil {
		return err
	}
	if liked {
		// remove from like
		err = utils.ReactionUpdator(data, a.store.RemoveFromLike)
		if err != nil {
			return err
		}
	}
	// update
	err = utils.ReactionUpdator(data, a.store.DislikeArticle)
	if err != nil {
		return err
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) AddCommentToArticle(data *NewComment) *types.CustomError {
	var wg sync.WaitGroup
	// check user data is complete
	errChan := make(chan *types.CustomError, 10)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := a.store.CheckUserProvidedData(data.UserID); err != nil {
			errChan <- err
			return
		}
		errChan <- nil
	}()

	go func() {
		wg.Wait()
		close(errChan)
	}()

	err := <-errChan
	if err != nil {
		return err
	}

	// check article exists
	exists, err := utils.CheckExistence(data.ArticleID, a.store.CheckArticleExist, 10)
	if err != nil {
		return err
	}
	if !exists {
		return utils.NewError(constants.NotFound, http.StatusNotFound)
	}

	// add new comment
	addErrChan := make(chan *types.CustomError, 10)
	go func() {
		if err := a.store.AddCommentToArticle(data); err != nil {
			addErrChan <- err
		}
		addErrChan <- nil
	}()

	err = <-addErrChan
	return err
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) EditArticleComment(data *EditComment) *types.CustomError {
	errChan := make(chan *types.CustomError, 1)
	go func() {
		err := a.store.EditArticleComment(data)
		errChan <- err
	}()

	err := <-errChan
	return err
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) EditArticleCommentByAdmin(data *EditCommentByAdmin) *types.CustomError {
	errChan := make(chan *types.CustomError, 1)
	go func() {
		err := a.store.EditArticleCommentByAdmin(data)
		errChan <- err
	}()

	err := <-errChan
	return err
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DeleteArticleComment(data *DeleteComment) *types.CustomError {
	errChan := make(chan *types.CustomError, 1)
	go func() {
		err := a.store.DeleteArticleComment(data)
		errChan <- err
	}()

	err := <-errChan
	return err
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) DeleteArticleCommentByAdmin(id uint) *types.CustomError {
	errChan := make(chan *types.CustomError, 1)
	go func() {
		err := a.store.DeleteArticleCommentByAdmin(id)
		errChan <- err
	}()

	err := <-errChan
	return err
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleService) ChangeStatus(status string, id uint) *types.CustomError {
	errChan := make(chan *types.CustomError, 1)
	go func() {
		err := a.store.ChangeStatus(status, id)
		errChan <- err
	}()
	err := <-errChan
	return err
}
