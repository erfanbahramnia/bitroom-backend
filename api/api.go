package api

import (
	"bitroom/article"
	"bitroom/auth"
	"bitroom/category"
	"bitroom/developer"
	"bitroom/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitApi(db *gorm.DB, ech *echo.Echo, port string) {
	// auth apis
	authStore := auth.NewAuthStore(db)
	authService := auth.NewAuthService(authStore)
	authHandler := auth.NewAuthHandler(authService)
	authHandler.InitHandler(ech)

	// user apis
	userStore := user.NewUserStore(db)
	userService := user.NewUserSerivce(userStore)
	userHandler := user.NewUserHandler(userService)
	userHandler.InitHandler(ech)

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

	// developer apis
	developer := developer.NewDeveloperApi(db, ech)
	developer.InitApi()

	// start server
	ech.Logger.Fatal(ech.Start(port))
}
