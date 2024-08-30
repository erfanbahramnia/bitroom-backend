package db

import (
	"bitroom/config"
	article_model "bitroom/models/article"
	category_model "bitroom/models/category"
	course_model "bitroom/models/course"
	news_model "bitroom/models/news"
	user_model "bitroom/models/user"
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

	db.AutoMigrate(
		&user_model.User{},
		&article_model.Article{},
		&article_model.ArticleProperty{},
		&article_model.ArticleComment{},
		&news_model.News{},
		&course_model.Course{},
		&category_model.Category{},
	)
	// success
	return db
}
