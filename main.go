package main

import (
	"Go_Project/controllers"
	_ "Go_Project/docs" // import swagger docs
	"Go_Project/middleware"
	"Go_Project/models"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go + Gin Book đào API
// @version 1.0
// @description This is a sample api.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email congdeptrai@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {

	models.ConnectDatabase()

	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/auth", controllers.GetAuth)
	api_v1 := r.Group("/api/v1")
	api_v1.Use(middleware.JWT())
	{
		api_v1.POST("/create-book/", controllers.CreateBook)
		api_v1.GET("/get-all-books/", controllers.GetAllBooks)
		api_v1.GET("/get-book-by-id/:id", controllers.GetBooksById)
	}
	r.Run(":8081")
}
