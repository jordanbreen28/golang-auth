package model

import (
	"api/database"
	"html"
	"strings"

	"api/utils/token"

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

func (u *User) UpdateUserDetails(updatedUser *User) (*User, error) {
	err := database.Database.Save(&updatedUser).Error
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

	err = database.Database.Model(User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (user *User) DeleteUser(id int) (User, error) {
	err := database.Database.Delete(&user, id).Error

	if err != nil {
		return User{}, err
	}
	return User{}, nil
}
