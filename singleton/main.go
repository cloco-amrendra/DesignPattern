// main.go
package main

import (
	"go-design-patterns/singleton/config"
	"go-design-patterns/singleton/handler"
	"go-design-patterns/singleton/middleware"
	"go-design-patterns/singleton/repository"
	"go-design-patterns/singleton/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database connection
	db := config.GetDB() // This  shows the singleton pattern use as the GetDB function is called only once

	// Set up repository, usecase, and handler layers
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	r := gin.Default()

	username := "admin"
	password := "password"

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(username, password))
	{
		protected.GET("/users", userHandler.GetUsers)
		protected.POST("/users", userHandler.CreateUser)
	}

	// Start the server
	r.Run(":8080")
}
