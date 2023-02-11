package models

type UserResponse struct {
	User_uuid string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
}
