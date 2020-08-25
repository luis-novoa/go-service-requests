package connection

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
  db, errors := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
	
	for _, err := range errors {
		fmt.Printf("errors: %v", err)
	}
	return db
}