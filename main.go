package main

import (
	"fmt"
	"gin/controller"
	"gin/entity"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.GET("/data", controller.ListUser)
	r.GET("/data/:id", controller.GetDate)
	r.Run()
	// Print a welcome message
	fmt.Println("Welcome to the Go Gin Gorm Web Service!")

	// Initialize the web service (this is just a placeholder)

}
