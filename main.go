package main

import (
	"api/controllers"
	"api/database"
	"api/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()

	router := gin.Default()

	superGroup := router.Group("/api/v1")
	{
		userGroup := superGroup.Group("/users")
		{
			userGroup.POST("/login", controllers.LoginUser)
			// new `GET /users` route associated with our `getUsers` function
			userGroup.GET("/", controllers.GetAllUsers)
			userGroup.POST("/", controllers.RegisterUser)
		}

		// todoGroup := superGroup.Group("/todo")
		// {
		// 	// todo group handlers
		// }
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
