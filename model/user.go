package model

import (
	"api/database"
	"fmt"
	"html"
	"strings"

	"api/utils/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Age      int    `gorm:"not null" json:"age"`
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) UpdateUserDetails(updatedUser *User, input UpdateUser) (*User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}
	input.Password = string(passwordHash)
	err = database.Database.Model(&updatedUser).Updates(input).Error
	if err != nil {
		return &User{}, err
	}
	return updatedUser, nil
}

func GetAllUsers() ([]User, error) {

	var users []User
	err := database.Database.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int) (User, error) {
	var user User
	err := database.Database.Find(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) Login(username string, password string) (string, error) {

	var err error

	u := User{}
	// find the user with the username
	err = database.Database.Model(User{}).Where("username = ?", username).Take(&u).Error

	// check if the user exists
	if err != nil {
		return "", err
	}
	// verify the password
	err = VerifyPassword(password, u.Password)
	// check if the password is correct
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	// generate a jwt token
	token, err := token.GenerateToken(u.ID)
	// check if there was an error generating the token
	if err != nil {
		return "", err
	}
	// return the token
	return token, nil
}

func (user *User) DeleteUser(id int, c *gin.Context) (User, error) {
	var tokenId uint
	var err error
	// extract the token id from the request
	tokenId, err = token.ExtractTokenID(c)
	// check if the token id is the same as the user id to be deleted
	if tokenId != uint(id) {
		return User{}, fmt.Errorf("You are not authorized to delete this user.")
	}
	// check if the token id is valid
	if err != nil || tokenId == 0 {
		return User{}, err
	}
	// Permenately delete the user at request
	err = database.Database.Unscoped().Delete(&user, id).Error
	// check if there was an error deleting the user
	if err != nil {
		return User{}, err
	}
	// return an empty user and nil error
	return User{}, nil
}
