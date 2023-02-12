package apiresponse

func ResponData(message string, code int, data interface{}) ResponseDt {
	jsonResp := ResponseDt{
		Status_code: code,
		Message:     message,
		Data:        data,
	}
	return jsonResp
}

func ResponStatus(message interface{}, code int) ResponseSts {
	jsonResp := ResponseSts{
		Status_code: code,
		Message:     message,
	}
	return jsonResp
}

func UserFormat(user_id string, name string, email string) UserResponse {
	jsonResp := UserResponse{
		User_id: user_id,
		Name:    name,
		Email:   email,
	}
	return jsonResp
}
