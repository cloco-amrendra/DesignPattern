package main

import (
	"go-design-patterns/singleton/config"
	"go-design-patterns/singleton/handler"
	"go-design-patterns/singleton/logger"
	"go-design-patterns/singleton/middleware"
	"go-design-patterns/singleton/repository"
	"go-design-patterns/singleton/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Get a single logger instance
	log := logger.GetLogger()
	log.Info("Starting application...")

	// Initialize the database connection
	db := config.GetDB()
	log.Info("Database connection established.")

	// Set up repository, usecase, and handler layers
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	r := gin.Default()

	// Load username and password from environment or configuration
	username := "admin"
	password := "password"

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
		log.Info("Ping endpoint was called.")
	})

	// Protected routes with basic authentication
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware(username, password))
	{
		protected.GET("/users", func(c *gin.Context) {
			log.Info("Fetching users")
			userHandler.GetUsers(c)
		})
		protected.POST("/users", func(c *gin.Context) {
			log.Info("Creating a new user")
			userHandler.CreateUser(c)
		})
	}

	log.Info("Server is running on port 8080.")
	// Start the server
	r.Run(":8080")
}
