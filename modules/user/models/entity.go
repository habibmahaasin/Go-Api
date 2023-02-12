package models

type User struct {
	User_id  string `json:"user_id"`
	Name     string `json:"name" binding:"required,min=4"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
