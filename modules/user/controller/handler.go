package controller

import (
	"errors"
	"gop-api/modules/user/models"
	"gop-api/modules/user/service"
	apiresponse "gop-api/package/api-response"
	bindvalidator "gop-api/package/bind-validator"
	jwttoken "gop-api/package/jwt-token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	service service.UserService
	token   jwttoken.JwtToken
}

func NewUserHandler(service service.UserService, token jwttoken.JwtToken) *userHandler {
	return &userHandler{service, token}
}

func (h *userHandler) User(c *gin.Context) {
	user, err := h.service.GetUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
		return
	}

	var userResponse []apiresponse.UserResponse
	for _, r := range user {
		response := apiresponse.UserFormat(r.User_id, r.Name, r.Email)
		userResponse = append(userResponse, apiresponse.UserResponse(response))
	}

	c.JSON(http.StatusOK, apiresponse.ResponData("Success", http.StatusOK, userResponse))
}

func (h *userHandler) AddUser(c *gin.Context) {
	var input models.AddUser

	if err := c.ShouldBindJSON(&input); err != nil {
		var validator validator.ValidationErrors
		if errors.As(err, &validator) {
			res := make([]bindvalidator.ErrorStruct, len(validator))
			for i, e := range validator {
				res[i] = bindvalidator.ErrorStruct{
					Field: e.Field(),
					Error: bindvalidator.BindingValidator(e),
				}
			}
			c.JSON(400, apiresponse.ResponStatus(res, http.StatusBadRequest))
		}
		return
	}

	checkEmail, _ := h.service.GetUserByEmail(input.Email)
	if checkEmail.Email != "" {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus("Email Already Used", http.StatusBadRequest))
		return
	}

	err := h.service.AddUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, apiresponse.ResponStatus("Success", http.StatusOK))
}

func (h *userHandler) Login(c *gin.Context) {
	var inputLogin models.InputLogin

	if err := c.ShouldBindJSON(&inputLogin); err != nil {
		var validator validator.ValidationErrors
		if errors.As(err, &validator) {
			res := make([]bindvalidator.ErrorStruct, len(validator))
			for i, e := range validator {
				res[i] = bindvalidator.ErrorStruct{
					Field: e.Field(),
					Error: bindvalidator.BindingValidator(e),
				}
			}
			c.JSON(400, apiresponse.ResponStatus(res, http.StatusBadRequest))
		}
		return
	}

	user, err := h.service.Login(inputLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apiresponse.ResponStatus(err.Error(), http.StatusUnauthorized))
		return
	}

	token, _ := h.token.GenerateJwt(user.User_id, user.Name)
	response := apiresponse.UserFormat(user.User_id, user.Name, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    response,
		"token":   token,
	})

}

func (h *userHandler) DetailUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
		return
	}

	if user.User_id == "" {
		c.JSON(http.StatusNotFound, apiresponse.ResponStatus("User Not Found", http.StatusNotFound))
		return
	}

	response := apiresponse.UserFormat(user.User_id, user.Name, user.Email)
	c.JSON(http.StatusOK, apiresponse.ResponData("Success", http.StatusOK, response))
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		var validator validator.ValidationErrors
		if errors.As(err, &validator) {
			res := make([]bindvalidator.ErrorStruct, len(validator))
			for i, e := range validator {
				res[i] = bindvalidator.ErrorStruct{
					Field: e.Field(),
					Error: bindvalidator.BindingValidator(e),
				}
			}
			c.JSON(400, apiresponse.ResponStatus(res, http.StatusBadRequest))
		}
		return
	}

	checkId, _ := h.service.GetUserById(id)
	if id != checkId.User_id {
		c.JSON(http.StatusNotFound, apiresponse.ResponStatus("User Not Found", http.StatusNotFound))
		return
	}

	checkEmail, _ := h.service.GetUserByEmail(input.Email)
	if checkEmail.Email != checkId.Email && checkEmail.Email != "" {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus("Email Already Used", http.StatusBadRequest))
		return
	}

	_, err := h.service.UpdateUser(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, apiresponse.ResponStatus("Success", http.StatusOK))
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	check, _ := h.service.GetUserById(id)
	if id != check.User_id {
		c.JSON(http.StatusNotFound, apiresponse.ResponStatus("User Not Found", http.StatusNotFound))
		return
	}

	_, err := h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, apiresponse.ResponStatus("Success", http.StatusOK))
}
