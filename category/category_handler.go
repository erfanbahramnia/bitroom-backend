package category

import (
	"bitroom/constants"
	"bitroom/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	service CategoryServiceInterface
}

func NewCategoryHandler(service CategoryServiceInterface) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (c *CategoryHandler) InitHandler(ech *echo.Echo) {
	group := ech.Group("category")

	group.POST("/add", c.AddCategory)
}

// @Summary Add Category
// @Description Adding new category by admin
// @Tags category
// @Accept json
// @Produce json
// @Param category body NewCategory true "Adding new category"
// @Success 201
// @Router /category/add [post]
func (c *CategoryHandler) AddCategory(ctx echo.Context) error {
	var category NewCategory

	// bind json to struct
	if err := ctx.Bind(&category); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constants.InvalidInputFormat)
	}

	// validate
	vs := utils.GetValidator()
	vsErrs := vs.Validate(category)
	if len(vsErrs) > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"errors": vsErrs,
		})
	}

	// add category
	err := c.service.AddCategory(category)
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}

	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": constants.CategoryAdded,
	})
}

func (c *CategoryHandler) GetCategories(ctx echo.Context) error {
	return nil
}

func (c *CategoryHandler) GetCategoriesTree(ctx echo.Context) error {
	return nil
}

func (c *CategoryHandler) GetCategoryById(ctx echo.Context) error {
	return nil
}

func (c *CategoryHandler) EditCategory(ctx echo.Context) error {
	return nil
}

func (c *CategoryHandler) DeleteCategory(ctx echo.Context) error {
	return nil
}
