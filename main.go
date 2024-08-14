package main

import (
	"bitroom/api"
	"bitroom/db"
)

func main() {
	// connect db
	db.InitDb()
	// start apis
	api.InitApi()
}
