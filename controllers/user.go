package controllers

import (
	"api/database"
	"api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /users
// Get all users
func GetAllUsers(c *gin.Context) {
	var users []model.User
	database.Database.Find(&users)

	c.JSON(http.StatusOK, users)
}

func RegisterUser(c *gin.Context) {
	var input model.RegistrationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Email:    input.Email,
		Age:      input.Age,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func LoginUser(c *gin.Context) {
	var input model.AuthenticationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{}

	user.Username = input.Username
	user.Password = input.Password

	token, err := user.Login(user.Username, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"JWT": token})
}
