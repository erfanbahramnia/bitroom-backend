package main

import (
	"bitroom/api"
	"bitroom/config"
	"bitroom/db"
	"bitroom/middleware"
	"fmt"

	_ "bitroom/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	// connect db
	db := db.InitDb()
	// configure server
	ech := echo.New()

	ech.Use(middleware.SetCsrfTokenMiddleware) // implement csrf

	ech.Static("/", "public") // static folder

	ech.GET("/api/*", echoSwagger.WrapHandler) // add swagger

	port := fmt.Sprintf(":%s", config.ServerConfig.Port)
	api.InitApi(db, ech, port)
}

// export GOOSE_DRIVER=postgres
// export GOOSE_DBSTRING=postgres://erfan:erfan.81@localhost:5432/bitroom_db?sslmode=disable
