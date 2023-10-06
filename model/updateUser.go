package model

type UpdateUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `gorm:"size:100;not null;unique" json:"email"`
	Age      int    `gorm:"not null" json:"age"`
}
