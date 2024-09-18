package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializaDB() {
	dsn := "host=localhost user=postgres password=123 dbname=book_management port=5432 sslmode=disable TimeZone=Asia/Karachi"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	log.Println("Database connected")

}
