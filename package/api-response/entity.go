package apiresponse

type ResponseDt struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

type ResponseSts struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type UserResponse struct {
	User_id string `json:"user_id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
}
