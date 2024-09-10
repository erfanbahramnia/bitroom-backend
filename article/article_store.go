package article

import (
	"bitroom/constants"
	article_model "bitroom/models/article"
	category_model "bitroom/models/category"
	user_model "bitroom/models/user"
	"bitroom/types"
	"bitroom/utils"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) *ArticleStore {
	return &ArticleStore{
		db: db,
	}
}

func (a *ArticleStore) GetCategory(id uint) (*category_model.Category, *types.CustomError) {
	// get category
	var category category_model.Category
	if err := a.db.Find(&category, id).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return &category, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) AddArticle(data *NewArticle) (*Article, *types.CustomError) {
	// get category
	category, err := a.GetCategory(data.Category)
	if err != nil {
		return nil, err
	}
	// save article
	article := article_model.Article{
		Title:       data.Title,
		Description: data.Description,
		Summary:     data.Summary,
		Image:       data.Image,
		Category:    *category,
	}
	if err := a.db.Create(&article).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return &Article{
		Title:       article.Title,
		Description: article.Description,
		Summary:     article.Summary,
		Image:       article.Image,
		ID:          article.ID,
		Status:      article.Status,
	}, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetArticles() ([]MinimumArticle, *types.CustomError) {
	var articles []MinimumArticle
	if err := a.db.Model(&article_model.Article{}).Select("title, summary, image, id").Where("status = ?", constants.ValidStatus[0]).Find(&articles).Error; err != nil {
		return nil, utils.NewError("could not retrive data", http.StatusInternalServerError)
	}
	// success
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetArticlesByAdmin() ([]MinimumArticle, *types.CustomError) {
	var articles []MinimumArticle
	if err := a.db.Model(&article_model.Article{}).Select("title, summary, image, id").Find(&articles).Error; err != nil {
		return nil, utils.NewError("could not retrive data", http.StatusInternalServerError)
	}
	// success
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetArticleById(id uint) (*article_model.Article, *types.CustomError) {

	// get article
	var article article_model.Article
	err := a.db.Preload("Properties").Preload("Comments").Preload("Category").First(&article, id).Error
	if err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	// update views counter
	article.Views++
	a.db.Save(&article)

	// success
	return &article, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetArticlesByCategory(categoryId uint) ([]MinimumArticle, *types.CustomError) {
	// recursive query for getting all children of selected category
	recursiveQuery := `
		WITH RECURSIVE category_tree as (
			SELECT * FROM categories WHERE id = ?
			UNION ALL
			SELECT c.* FROM categories c
			INNER JOIN category_tree ct ON ct.id = c.Parent_id
		)
		SELECT id FROM category_tree;
	`
	var ids []uint
	if err := a.db.Model(&category_model.Category{}).Raw(recursiveQuery, categoryId).Select("id").Scan(&ids).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	// get articles
	var articles []MinimumArticle
	if err := a.db.Model(&article_model.Article{}).
		Select("articles.title, articles.summary, articles.image, articles.id").
		Joins("LEFT JOIN categories ON categories.id = articles.category_id").
		Where("categories.id IN ?", ids).
		Scan(&articles).Error; err != nil {
		return nil, utils.NewError("could not retrive data", http.StatusInternalServerError)
	}
	// success
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) EditArticle(editedArticle *EditArticle) *types.CustomError {
	article := article_model.Article{}
	// get article
	if err := a.db.First(&article, *editedArticle.Id).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	// update
	if editedArticle.Title != nil {
		article.Title = *editedArticle.Title
	}
	if editedArticle.Description != nil {
		article.Description = *editedArticle.Description
	}
	if editedArticle.Summary != nil {
		article.Summary = *editedArticle.Summary
	}
	if editedArticle.Status != nil {
		article.Status = *editedArticle.Status
	}
	if editedArticle.Image != nil {
		article.Image = *editedArticle.Image
	}
	if editedArticle.Category != nil {
		// get category
		category, err := a.GetCategory(*editedArticle.Category)
		if err != nil {
			return err
		}
		article.Category = *category
	}

	// save changes
	if err := a.db.Save(&article).Error; err != nil {
		return utils.NewError("failed to update", http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DeleteArticle(id uint) *types.CustomError {
	// delete
	if err := a.db.Delete(&article_model.Article{}, id).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) AddArticleProperty(data *ArticleProperty) *types.CustomError {
	// add new property
	property := article_model.ArticleProperty{
		Image:       *data.Image,
		Description: data.Description,
		ArticleID:   data.ArticleID,
	}
	if err := a.db.Create(&property).Error; err != nil {
		return utils.NewError("could not add new property", http.StatusInternalServerError)
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) EditArticleProperty(data *EditArticleProperty) *types.CustomError {
	var property article_model.ArticleProperty
	if err := a.db.First(&property, data.PropertyID).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	// update
	if data.Description != nil {
		property.Description = *data.Description
	}
	if data.Image != nil {
		property.Image = *data.Image
	}

	// save changes
	if err := a.db.Save(property).Error; err != nil {
		return utils.NewError("failed to update", http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DeleteArticleProperty(id uint) *types.CustomError {
	if err := a.db.Delete(&article_model.ArticleProperty{}, id).Error; err != nil {
		return utils.NewError(constants.CouldNotDeleteItem, http.StatusInternalServerError)
	}
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) GetPopularArticles() ([]MinimumArticle, *types.CustomError) {
	var articles []MinimumArticle
	if err := a.db.Model(&article_model.Article{}).Order("views desc").Limit(10).Find(&articles).Error; err != nil {
		return nil, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return articles, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) LikeArticle(data *types.LikeOrDislikeArticle) *types.CustomError {
	if err := a.db.Model(&article_model.Article{}).
		Where("id = ?", data.ArticleId).
		Update("likes", gorm.Expr("array_append(likes, ?)", data.UserId)).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DislikeArticle(data *types.LikeOrDislikeArticle) *types.CustomError {
	if err := a.db.Model(&article_model.Article{}).
		Where("id = ?", data.ArticleId).
		Update("dislikes", gorm.Expr("array_append(dislikes, ?)", data.UserId)).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) AddCommentToArticle(data *NewComment) *types.CustomError {
	// create new comment
	comment := article_model.ArticleComment{
		Comment:   data.Comment,
		ArticleID: data.ArticleID,
		UserID:    data.UserID,
	}
	if err := a.db.Create(&comment).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) EditArticleComment(data *EditComment) *types.CustomError {
	// update comment
	result := a.db.Model(&article_model.ArticleComment{}).
		Where("article_id = ? AND user_id = ? AND id = ?", data.ArticleId, data.UserID, data.CommentId).
		Updates(&article_model.ArticleComment{Comment: data.Comment})

	if result.Error != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	if result.RowsAffected == 0 {
		return utils.NewError("comment not found or no changes made", http.StatusNotFound)
	}

	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) EditArticleCommentByAdmin(data *EditCommentByAdmin) *types.CustomError {
	// update comment
	result := a.db.Model(&article_model.ArticleComment{}).
		Where("id = ?", data.CommentId).
		Updates(&article_model.ArticleComment{Comment: data.Comment})

	if result.Error != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	if result.RowsAffected == 0 {
		return utils.NewError("comment not found or no changes made", http.StatusNotFound)
	}

	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DeleteArticleComment(data *DeleteComment) *types.CustomError {
	result := a.db.Where("article_id = ? AND user_id = ? AND id = ?", data.ArticleId, data.UserID, data.CommentId).
		Delete(&article_model.ArticleComment{})

	if result.Error != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	if result.RowsAffected == 0 {
		return utils.NewError("comment not found or no changes made", http.StatusNotFound)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) DeleteArticleCommentByAdmin(id uint) *types.CustomError {
	result := a.db.Where("id = ?", id).
		Delete(&article_model.ArticleComment{})

	if result.Error != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	if result.RowsAffected == 0 {
		return utils.NewError("comment not found or no changes made", http.StatusNotFound)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckArticleExist(id uint) (bool, *types.CustomError) {
	var exists bool
	err := a.db.Model(&article_model.Article{}).
		Select("count(*) > 0").
		Where("ID = ? AND status = ?", id, constants.ValidStatus[0]).
		Scan(&exists).Error

	if err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	return exists, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckArticleAllStatusExist(id uint) (bool, *types.CustomError) {
	var exists bool
	err := a.db.Model(&article_model.Article{}).
		Select("count(*) > 0").
		Where("ID = ?", id).
		Scan(&exists).Error

	if err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	return exists, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckCategoryExist(id uint) (bool, *types.CustomError) {
	var exists bool
	err := a.db.Model(&category_model.Category{}).Select("count(*) > 0").Where("ID = ?", id).Scan(&exists).Error
	if err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return exists, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckPropertyExists(id uint) (bool, *types.CustomError) {
	var exists bool
	err := a.db.Model(&article_model.ArticleProperty{}).
		Select("count(*) > 0").
		Where("ID = ?", id).
		Scan(&exists).Error
	if err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	return exists, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckUserDisliked(data *types.LikeOrDislikeArticle) (bool, *types.CustomError) {
	var count int64

	if err := a.db.Model(&article_model.Article{}).Where("id = ?", data.ArticleId).Where("? = ANY(dislikes)", data.UserId).Count(&count).Error; err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	exists := count > 0
	return exists, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) RemoveFromDislike(data *types.LikeOrDislikeArticle) *types.CustomError {
	if err := a.db.Model(&article_model.Article{}).
		Where("id = ?", data.ArticleId).
		Update("dislikes", gorm.Expr("array_remove(dislikes, ?)", data.UserId)).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckUserLiked(data *types.LikeOrDislikeArticle) (bool, *types.CustomError) {
	var count int64

	if err := a.db.Model(&article_model.Article{}).Where("id = ?", data.ArticleId).Where("? = ANY(likes)", data.UserId).Count(&count).Error; err != nil {
		return false, utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	exists := count > 0
	return exists, nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) RemoveFromLike(data *types.LikeOrDislikeArticle) *types.CustomError {
	if err := a.db.Model(&article_model.Article{}).
		Where("id = ?", data.ArticleId).
		Update("likes", gorm.Expr("array_remove(likes, ?)", data.UserId)).Error; err != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// success
	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) CheckUserProvidedData(userId uint) *types.CustomError {
	var user user_model.User
	// get user
	if err := a.db.First(&user, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewError("user", http.StatusNotFound)
		}
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}
	// check user data is complete
	if user.FirstName == "" || user.LastName == "" {
		return utils.NewError("incomplete user data", http.StatusBadRequest)
	}

	return nil
}

// --------------------------------------------------------------------------------------------------------------------

func (a *ArticleStore) ChangeStatus(status string, id uint) *types.CustomError {
	// update status
	result := a.db.Model(&article_model.Article{}).Where("id = ?", id).Updates(&article_model.Article{Status: status})

	if result.Error != nil {
		return utils.NewError(constants.InternalServerError, http.StatusInternalServerError)
	}

	if result.RowsAffected == 0 {
		return utils.NewError("comment not found or no changes made", http.StatusNotFound)
	}

	// success
	return nil
}
