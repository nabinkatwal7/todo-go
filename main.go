package main

import (
	"fmt"
	"log"
	"todo-api/controller"
	"todo-api/db"
	"todo-api/middleware"
	"todo-api/model"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	db.Connect()
	db.Database.AutoMigrate(&model.User{})
	db.Database.AutoMigrate(&model.ToDo{})
}

func serveApplication() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server responded with status code 200",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/todo", controller.AddTodo)
	protectedRoutes.GET("/todo", controller.GetAllEntries)
	protectedRoutes.PUT("/todo/:id", controller.UpdateTodo)
	protectedRoutes.DELETE("/todo/:id", controller.DeleteTodo)

	router.Run(":8080")
	fmt.Println("Server started on port 8080")
}