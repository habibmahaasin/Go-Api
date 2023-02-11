package controller

import (
	jwttoken "gop-api/app/jwt-token"
	"gop-api/modules/user/models"
	"gop-api/modules/user/service"
	apiresponse "gop-api/package/api-response"
	"net/http"

	"github.com/gin-gonic/gin"
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

	var userResponse []models.UserResponse

	for _, r := range user {
		response := apiresponse.ResponeFormat(r.User_uuid, r.Name, r.Email)
		userResponse = append(userResponse, models.UserResponse(response))
	}

	c.JSON(http.StatusOK, apiresponse.ResponData("Success", http.StatusOK, userResponse))
}

func (h *userHandler) AddUser(c *gin.Context) {
	var input models.AddUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
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
		response := apiresponse.ResponStatus(err.Error(), http.StatusBadRequest)
		c.JSON(400, response)
		return
	}

	user, err := h.service.Login(inputLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, apiresponse.ResponStatus(err.Error(), http.StatusUnauthorized))
		return
	}

	token, _ := h.token.GenerateJwt(user.User_uuid, user.Name)
	response := apiresponse.ResponeFormat(user.User_uuid, user.Name, user.Email)

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

	if user.User_uuid == "" {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus("User Not Found", http.StatusBadRequest))
		return
	}

	response := apiresponse.ResponeFormat(user.User_uuid, user.Name, user.Email)
	c.JSON(http.StatusOK, apiresponse.ResponData("Success", http.StatusOK, response))
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
		return
	}

	_, err := h.service.UpdateUser(id, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
		return
	}

	userUpdated, err := h.service.GetUserById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
		return
	}

	if userUpdated.User_uuid == "" {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus("User Not Found", http.StatusBadRequest))
		return
	}

	response := apiresponse.ResponeFormat(userUpdated.User_uuid, userUpdated.Name, userUpdated.Email)
	c.JSON(http.StatusOK, apiresponse.ResponData("Success", http.StatusOK, response))
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	check, _ := h.service.GetUserById(id)
	if id != check.User_uuid {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus("User Not Found", http.StatusBadRequest))
		return
	}

	_, err := h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, apiresponse.ResponStatus(err.Error(), http.StatusBadRequest))
		return
	}

	c.JSON(http.StatusOK, apiresponse.ResponStatus("Success", http.StatusOK))
}
