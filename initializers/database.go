package initializers

import (
	"cobacoba1/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:8111)/iventorybook"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}

	database.AutoMigrate(&models.User{})

	DB = database
}
