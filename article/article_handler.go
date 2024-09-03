package article

import (
	"bitroom/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	service ArticleServiceInterface
}

func NewArticleHandler(service ArticleServiceInterface) *ArticleHandler {
	return &ArticleHandler{
		service: service,
	}
}

func (a *ArticleHandler) InitHandler(ech *echo.Echo) {
	group := ech.Group("article")

	group.POST("/add", a.AddArticle)
	group.GET("/all", a.GetArticles)
	group.GET("/popular", a.GetPopularArticles)
	group.GET("/:id", a.GetArticleById)
	group.PUT("/edit", a.EditArticle)
	group.DELETE("/:id", a.DeleteArticleById)

	group.GET("/byCategory/:categoryId", a.GetArticlesByCategory)

	group.POST("/property/add", a.AddArticleProperty)
	group.PUT("/property/edit", a.EditArticleProperty)
	group.DELETE("/property/:propertyId", a.DeleteArticleProperty)
}

// AddArticle godoc
// @Description Upload an article along with an image
// @Tags articles
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Article Title"
// @Param description formData string true "Article Description"
// @Param summary formData string true "Article Summary"
// @Param category formData string true "Article Category"
// @Param image formData file true "Article Image"
// @Router /article/add [post]
func (a *ArticleHandler) AddArticle(ctx echo.Context) error {
	// get data
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	summary := ctx.FormValue("summary")
	// Validate required fields
	if title == "" || description == "" || summary == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "All fields are required")
	}

	// get category id
	category, ParsingErr := strconv.ParseUint(ctx.FormValue("category"), 10, 64)
	if ParsingErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please upload valid id")
	}

	// save file
	file, err := ctx.FormFile("image")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "upload image")
	}
	uploadPath, uploadErr := utils.HanldeFileUpload(file)
	if uploadErr != nil {
		return echo.NewHTTPError(uploadErr.Code, uploadErr.Message)
	}

	// save article
	article := &NewArticle{
		Title:       title,
		Description: description,
		Summary:     summary,
		Category:    uint(category),
		Image:       uploadPath,
	}
	InsertedData, addErr := a.service.AddArticle(article)
	if addErr != nil {
		return echo.NewHTTPError(addErr.Code, addErr.Message)
	}

	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"article": InsertedData,
	})
}

// --------------------------------------------------------------------------------------------------------------------

// @Description get all articles
// @Tags articles
// @Produce json
// @Router /article/all [get]
func (a *ArticleHandler) GetArticles(ctx echo.Context) error {
	articles, err := a.service.GetArticles()
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"articles": articles,
	})
}

// --------------------------------------------------------------------------------------------------------------------

// @Summary Get Article By Id
// @Tags articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 201
// @Router /article/{id} [get]
func (a *ArticleHandler) GetArticleById(ctx echo.Context) error {
	id, ParsingErr := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if ParsingErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please upload valid id")
	}

	article, err := a.service.GetArticleById(uint(id))
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"article": article,
	})
}

// --------------------------------------------------------------------------------------------------------------------

// @Summary Delete Article By Id
// @Tags articles
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Success 201
// @Router /article/{id} [delete]
func (a *ArticleHandler) DeleteArticleById(ctx echo.Context) error {
	// get article id
	id, ParsingErr := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if ParsingErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please upload valid id")
	}
	// delete
	err := a.service.DeleteArticle(uint(id))
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}
	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "successfully deleted",
	})
}

// --------------------------------------------------------------------------------------------------------------------

// @Summary Get Articles By category
// @Tags articles
// @Accept json
// @Produce json
// @Param categoryId path int true "Category ID"
// @Success 201
// @Router /article/byCategory/{categoryId} [get]
func (a *ArticleHandler) GetArticlesByCategory(ctx echo.Context) error {
	categoryId, ParsingErr := strconv.ParseUint(ctx.Param("categoryId"), 10, 64)
	if ParsingErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please upload valid id")
	}
	// get articles
	articles, err := a.service.GetArticlesByCategory(uint(categoryId))
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}
	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"articles": articles,
	})
}

// --------------------------------------------------------------------------------------------------------------------

// @Description Edit an article
// @Tags articles
// @Accept multipart/form-data
// @Produce json
// @Param title formData string false "Article Title"
// @Param description formData string false "Article Description"
// @Param summary formData string false "Article Summary"
// @Param category formData string false "Article Category"
// @Param status formData string false "Article Status"
// @Param id formData string false "Article id"
// @Param image formData file false "Article Image"
// @Router /article/edit [put]
func (a *ArticleHandler) EditArticle(ctx echo.Context) error {
	var editedArticle EditArticle

	// get article id
	// Bind form data to struct
	if err := ctx.Bind(&editedArticle); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	// save file if sended
	if file, err := ctx.FormFile("image"); file != nil {
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to get image")
		}
		uploadPath, uploadErr := utils.HanldeFileUpload(file)
		if uploadErr != nil {
			return echo.NewHTTPError(uploadErr.Code, uploadErr.Message)
		}
		editedArticle.Image = &uploadPath
	}

	// check any data sended for update
	if editedArticle.Title == nil &&
		editedArticle.Description == nil &&
		editedArticle.Summary == nil &&
		editedArticle.Status == nil &&
		editedArticle.Category == nil &&
		editedArticle.Image == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please upload any data for update")
	}

	// update
	err := a.service.EditArticle(&editedArticle)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "updated successfully",
	})
}

// --------------------------------------------------------------------------------------------------------------------

// @Tags articles
// Accept multipart/form-data
// @Product json
// @Param description formData string false "Property description"
// @Param article_id formData uint true "Article id"
// @Param image formData file false "Property image"
// @Router /article/property/add [post]
func (a *ArticleHandler) AddArticleProperty(ctx echo.Context) error {
	var property ArticleProperty
	// bind json to struct
	if err := ctx.Bind(&property); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	// validate
	vs := utils.GetValidator()
	vsErr := vs.Validate(property)
	if len(vsErr) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"errors": vsErr,
		})
	}

	// save file
	file, fileErr := ctx.FormFile("image")
	if fileErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "upload image")
	}
	uploadPath, err := utils.HanldeFileUpload(file)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}
	property.Image = &uploadPath

	// add property
	err = a.service.AddArticleProperty(&property)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}

// --------------------------------------------------------------------------------------------------------------------

// @Tags articles
// Accept multipart/form-data
// @Product json
// @Param description formData string false "Property description"
// @Param property_id formData uint true "Article id"
// @Param image formData file false "Property image"
// @Router /article/property/edit [put]
func (a *ArticleHandler) EditArticleProperty(ctx echo.Context) error {
	var propertyEdit EditArticleProperty
	// bind form data to struct
	if err := ctx.Bind(&propertyEdit); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"message": "Invalid input",
			"error":   err.Error(),
		})
	}

	// save file if has been sended
	if file, err := ctx.FormFile("image"); file != nil {
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "failed to get image")
		}
		uploadPath, uploadErr := utils.HanldeFileUpload(file)
		if uploadErr != nil {
			return echo.NewHTTPError(uploadErr.Code, uploadErr.Message)
		}
		propertyEdit.Image = &uploadPath
	}

	// check any data have been sended or not
	if propertyEdit.Description == nil && propertyEdit.Image == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please upload any data for update")
	}

	// update
	err := a.service.EditArticleProperty(&propertyEdit)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}

// --------------------------------------------------------------------------------------------------------------------

// @Tags articles
// Accept multipart/form-data
// @Product json
// @Param propertyId path int true "Property ID"
// @Router /article/property/{propertyId} [delete]
func (a *ArticleHandler) DeleteArticleProperty(ctx echo.Context) error {
	// get id of property
	propertyId, parsingErr := strconv.ParseUint(ctx.Param("propertyId"), 10, 64)
	if parsingErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "please upload valid id")
	}
	// delete
	err := a.service.DeleteArticleProperty(uint(propertyId))
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "ok",
	})
}

// --------------------------------------------------------------------------------------------------------------------

// @Description get all popular articles
// @Tags articles
// @Produce json
// @Router /article/popular [get]
func (a *ArticleHandler) GetPopularArticles(ctx echo.Context) error {

	// get articles
	articles, err := a.service.GetPopularArticles()
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"articles": articles,
	})
}
