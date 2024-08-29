package api

import (
	"bitroom/article"
	"bitroom/auth"
	"bitroom/category"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Swagger Example API
// @version 1.0
// @BasePath /

func InitApi(db *gorm.DB, ech *echo.Echo, port string) {
	// auth apis
	authStore := auth.NewAuthStore(db)
	authService := auth.NewAuthService(authStore)
	authHandler := auth.NewAuthHandler(authService)
	authHandler.InitHandler(ech)

	// category apis
	categoryStore := category.NewCategoryStore(db)
	categoryService := category.NewCategoryService(categoryStore)
	categoryHandler := category.NewCategoryHandler(categoryService)
	categoryHandler.InitHandler(ech)

	// article apis
	articleStore := article.NewArticleStore(db)
	articleService := article.NewArticleService(articleStore)
	articleHandler := article.NewArticleHandler(articleService)
	articleHandler.InitHandler(ech)

	// start server
	ech.Logger.Fatal(ech.Start(port))
}
