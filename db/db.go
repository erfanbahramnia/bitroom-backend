package db

import (
	"bitroom/config"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	// connecting to database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbConfig.DB_HOST, config.DbConfig.DB_PORT, config.DbConfig.DB_USER, config.DbConfig.DB_PASS, config.DbConfig.DB_NAME)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		log.Fatal("could not connect db...")
	}

	// success
	return db
}
