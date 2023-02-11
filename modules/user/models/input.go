package models

import "github.com/google/uuid"

type AddUser struct {
	User_uuid uuid.UUID
	Name      string `json:"name" binding:"required,min=4"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
}

type InputLogin struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
