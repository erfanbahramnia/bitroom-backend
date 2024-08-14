package db

import (
	"bitroom/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDb() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbConfig.DB_HOST, config.DbConfig.DB_PORT, config.DbConfig.DB_USER, config.DbConfig.DB_PASS, config.DbConfig.DB_NAME)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		log.Fatal("could not connect db...")
	}

	getPing(db)

	return db
}

func getPing(db *sql.DB) {
	if err := db.Ping(); err != nil {
		fmt.Println(err)
		log.Fatal("could not get ping from db...")
	}
}
