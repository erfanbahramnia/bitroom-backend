package article

import (
	"bitroom/utils"
	"net/http"

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
}

// AddArticle godoc
// @Description Upload an article along with an image
// @Tags articles
// @Accept multipart/form-data
// @Produce json
// @Param title formData string true "Article Title"
// @Param description formData string true "Article Description"
// @Param summary formData string true "Article Summary"
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
