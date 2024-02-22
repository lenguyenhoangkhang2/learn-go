package database

import (
	"log"

	"github.com/lenguyenhoangkhang2/go_authentication/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connect to Database!")
}

func Migrate() {
	Instance.AutoMigrate(models.User{})

	log.Println("Database migration completed!")
}
