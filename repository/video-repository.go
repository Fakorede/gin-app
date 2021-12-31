package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/Fakorede/gin-app/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Save(video entity.Video)
	Update(video entity.Video)
	Delete(video entity.Video)
	FindAll() []entity.Video
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepository() VideoRepository {
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_database := os.Getenv("DB_DATABASE")

	if db_host == "" || db_port == "" || db_database == "" || db_user == "" || db_pass == "" {
		log.Fatal("Database connection variables required in .env")
	}

	dsn := db_user + ":" + db_pass + "@tcp(" + db_host + ":" + db_port + ")/" + db_database + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to db")
	}

	db.AutoMigrate(
		&entity.Video{},
		&entity.Person{},
	)

	exists := db.Migrator().HasTable(&entity.Person{})
	log.Println(fmt.Sprint(exists) + " => it does or not")

	return &database{
		connection: db,
	}
}

func (db *database) Save(video entity.Video) {
	db.connection.Create(&video)
}

func (db *database) Update(video entity.Video) {
	db.connection.Save(&video)
}

func (db *database) Delete(video entity.Video) {
	db.connection.Delete(&video)
}

func (db *database) FindAll() []entity.Video {
	var videos []entity.Video
	db.connection.Preload("Author").Find(&videos)
	return videos
}

func (db *database) CloseDB() {
	sqlDB := db.connection
	conn, err := sqlDB.DB()
	if err != nil {
		panic("Failed to connect to db")
	}

	conn.Close()
}
