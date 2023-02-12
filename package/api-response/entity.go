package apiresponse

type ResponseDt struct {
	Message     string      `json:"message"`
	Status_code int         `json:"status_code"`
	Data        interface{} `json:"data"`
}

type ResponseSts struct {
	Message     interface{} `json:"message"`
	Status_code int         `json:"status_code"`
}

type UserResponse struct {
	User_id string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}
