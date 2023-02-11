package apiresponse

func ResponData(message string, code int, data interface{}) ResponseDt {
	jsonResp := ResponseDt{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return jsonResp
}

func ResponStatus(message string, code int) ResponseSts {
	jsonResp := ResponseSts{
		Code:    code,
		Message: message,
	}
	return jsonResp
}

func ResponeFormat(user_uuid string, name string, email string) UserResponse {
	jsonResp := UserResponse{
		User_uuid: user_uuid,
		Name:      name,
		Email:     email,
	}
	return jsonResp
}
