package api

import (
	"bitroom/auth"
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// @title Swagger Example API
// @version 1.0
// @BasePath /

func InitApi(db *gorm.DB, ech *echo.Echo, port string) {
	fmt.Println("api")
	authStore := auth.NewAuthStore(db)
	authService := auth.NewAuthService(authStore)
	authHandler := auth.NewAuthHandler(authService)
	authHandler.InitHandler(ech)

	ech.Logger.Fatal(ech.Start(port))
}
