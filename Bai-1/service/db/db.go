package db

import (
	"awesomeProject6/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg *config.Config) {
	dsn := fmt.Sprint("host=127.0.0.1 user=postgres password=password dbname=movie_service port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fail to connect db: ", err)
	}

	fmt.Println("Connected db")
}
