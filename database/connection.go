package database

import (
	"os"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func connect() {
	const (
		host = "localhost"
		port = 5432
		user = os.Getenv("DB_ROLE")
		password = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, user, password, dbname)
  db, errors := gorm.Open("postgres", psqlInfo)
	
	for _, err := range errors {
		fmt.Printf("errors: %v", err)
	}
	return db
}