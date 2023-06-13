package main

import (
	"cobacoba1/initializers"
	"cobacoba1/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
