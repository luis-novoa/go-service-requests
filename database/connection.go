package database

import (
	"os"
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/luis-novoa/go-service-requests/models"
)

func Connect() *gorm.DB {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
	}
	
	host := "localhost"
	port := 5432
	user := os.Getenv("DB_ROLE")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, errors := gorm.Open("postgres", psqlInfo)
	
	db.AutoMigrate(&models.User{}, &models.ServiceRequest{})
	
	if errors != nil {
		fmt.Printf("errors: %v", errors)
	}
	return db
}