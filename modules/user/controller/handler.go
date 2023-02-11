package controller

import (
	"gop-api/modules/user/models"
	"gop-api/modules/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *userHandler {
	return &userHandler{service}
}

func (h *userHandler) User(c *gin.Context) {
	user, err := h.service.GetUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "User Handler Error",
		})
		return
	}

	var userResponse []models.UserResponse

	for _, r := range user {
		response := models.UserResponse{
			User_uuid: r.User_uuid,
			Email:     r.Email,
			Name:      r.Name,
		}
		userResponse = append(userResponse, response)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    userResponse,
	})
}

func (h *userHandler) AddUser(c *gin.Context) {
	var input models.AddUser

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	err := h.service.AddUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}

func (h *userHandler) Login(c *gin.Context) {
	var inputLogin models.InputLogin

	if err := c.ShouldBindJSON(&inputLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	user, err := h.service.Login(inputLogin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	response := models.UserResponse{
		User_uuid: user.User_uuid,
		Email:     user.Email,
		Name:      user.Name,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    response,
	})

}

func (h *userHandler) DetailUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		})
		return
	}

	response := models.UserResponse{
		User_uuid: user.User_uuid,
		Email:     user.Email,
		Name:      user.Name,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"detail":  response,
	})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	_, err := h.service.UpdateUser(id, input)
	userUpdated, err := h.service.GetUserById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	response := models.UserResponse{
		User_uuid: userUpdated.User_uuid,
		Email:     userUpdated.Email,
		Name:      userUpdated.Name,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":       http.StatusOK,
		"message":      "Success",
		"data_updated": response,
	})
}

func (h *userHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	_, err := h.service.DeleteUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}
