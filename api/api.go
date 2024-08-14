package api

import (
	"bitroom/config"
	"fmt"

	"github.com/labstack/echo/v4"
)

func InitApi() {
	e := echo.New()

	port := fmt.Sprintf(":%s", config.ServerConfig.Port)
	fmt.Println(port)
	e.Logger.Fatal(e.Start(port))
}
