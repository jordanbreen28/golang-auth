package main

import (
	"api/controllers"
	"api/database"
	"api/middleware"
	"api/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	loadEnv()
	loadDatabase()
}

func main() {
	router := gin.Default()

	superGroup := router.Group("/api/v1")
	{
		userGroup := superGroup.Group("/users")
		{
			userGroup.POST("/", controllers.RegisterUser)
			userGroup.POST("/login", controllers.LoginUser)
			// below routes are protected by jwt auth middleware
			userGroup.Use(middleware.JwtAuthMiddleware())
			userGroup.POST("/logout", controllers.Logout)
			userGroup.GET("/", controllers.GetAllUsers)
			userGroup.GET("/:id", controllers.GetUserById)
			userGroup.PUT("/:id", controllers.UpdateUser)
			userGroup.DELETE("/:id", controllers.DeleteUser)
		}
	}
	router.Run("0.0.0.0:8080")
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
