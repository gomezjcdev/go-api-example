package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DSN = "host=localhost user=gomezjc password='' dbname=go_example port=5432"
var DB *gorm.DB

func DBConnection() {
	var dbError error
	DB, dbError = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
	} else {
		log.Println("DB connected")
	}
}
