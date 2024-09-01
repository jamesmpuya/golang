package config

import (
	"log"
	"os"

	"github.com/jamesmpuya/tapindigital/tapinapi-gin/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("POSTGRES_DB")), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}
