package main

import (
	"cobacoba1/controllers"
	"cobacoba1/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	v1 := r.Group("v1")
	v1.POST("/user", controllers.CreateUser)
	v1.GET("/users", controllers.GetAllUser)
	v1.GET("/user/:id", controllers.GetUserById)
	v1.POST("/login", controllers.LoginUser)

	r.Run()
}
