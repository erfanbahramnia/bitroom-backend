package category

import (
	"bitroom/constants"
	"bitroom/middleware"
	"bitroom/utils"
	"net/http"
	"strconv"

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

	group.POST("/add", c.AddCategory, middleware.JwtMiddleware, middleware.RoleBaseMiddleware([]string{"admin", "administrator"}))
	group.GET("/:id", c.GetCategoryById)
	group.GET("/all", c.GetCategories)
	group.GET("/tree", c.GetCategoriesTree)
	group.PUT("/:id/:name", c.EditCategory, middleware.JwtMiddleware, middleware.RoleBaseMiddleware([]string{"admin", "administrator"}))
	group.DELETE("/:id", c.DeleteCategory, middleware.JwtMiddleware, middleware.RoleBaseMiddleware([]string{"admin", "administrator"}))
}

// @Summary Add Category
// @Description Adding new category by admin
// @Tags category
// @Accept json
// @Produce json
// @Param category body NewCategory true "Adding new category"
// @Success 201
// @Router /category/add [post]
// @Security BearerAuth
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

// --------------------------------------------------------------------------------------------------------

// @Summary Get Categories
// @Tags category
// @Accept json
// @Produce json
// @Success 201
// @Router /category/all [get]
func (c *CategoryHandler) GetCategories(ctx echo.Context) error {
	categories, err := c.service.GetCategories()
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}
	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"categories": categories,
	})
}

// --------------------------------------------------------------------------------------------------------

// @Summary Get Categorys with tree format
// @Tags category
// @Accept json
// @Produce json
// @Success 201
// @Router /category/tree [get]
func (c *CategoryHandler) GetCategoriesTree(ctx echo.Context) error {
	categories, err := c.service.GetCategoriesTree()
	if err != nil {
		return echo.NewHTTPError(err.Code, err.Message)
	}
	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"categories": categories,
	})
}

// --------------------------------------------------------------------------------------------------------

// @Summary Get Category By Id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 201
// @Router /category/{id} [get]
func (c *CategoryHandler) GetCategoryById(ctx echo.Context) error {
	// get category id
	categoryId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil || categoryId < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, constants.ProvideId)
	}
	// get category
	category, getErr := c.service.GetCategoryById(uint(categoryId))
	if getErr != nil {
		return echo.NewHTTPError(getErr.Code, getErr.Message)
	}
	// success
	return ctx.JSON(http.StatusOK, category)
}

// --------------------------------------------------------------------------------------------------------

// @Summary Edit Category By Id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param name path string true "New Name"
// @Success 201
// @Router /category/{id}/{name} [put]
// @Security BearerAuth
func (c *CategoryHandler) EditCategory(ctx echo.Context) error {
	// get new name
	name := ctx.Param("name")
	if len(name) < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, constants.ProvideId)
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil || id < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, constants.ProvideId)
	}
	data := EditCategory{
		Name: name,
		ID:   uint(id),
	}
	// update
	editErr := c.service.EditCategory(data)
	if editErr != nil {
		return echo.NewHTTPError(editErr.Code, editErr.Message)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "successful",
	})
}

// --------------------------------------------------------------------------------------------------------

// @Summary Delete Category By Id
// @Tags category
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 201
// @Router /category/{id} [delete]
// @Security BearerAuth
func (c *CategoryHandler) DeleteCategory(ctx echo.Context) error {
	// get id
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil || id < 1 {
		return echo.NewHTTPError(http.StatusBadRequest, constants.ProvideId)
	}
	// delete
	deleteErr := c.service.DeleteCategory(uint(id))
	if deleteErr != nil {
		return echo.NewHTTPError(deleteErr.Code, deleteErr.Message)
	}
	// success
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg": "successful",
	})
}
