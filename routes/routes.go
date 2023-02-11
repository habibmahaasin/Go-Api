package routes

import (
	jwttoken "gop-api/app/jwt-token"
	"gop-api/app/middlewares"
	"gop-api/modules/user/controller"
	"gop-api/modules/user/repository"
	"gop-api/modules/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	token := jwttoken.NewJwtToken()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := controller.NewUserHandler(userService, token)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "endpoint not found",
		})
	})

	apiV1 := router.Group("api/user/v1")
	apiV1.POST("/login", userHandler.Login)
	apiV1.POST("/create", userHandler.AddUser)
	apiV1.GET("/list", middlewares.ApiAuth(token, userService), userHandler.User)
	apiV1.GET("/detail/:id", middlewares.ApiAuth(token, userService), userHandler.DetailUser)
	apiV1.PUT("/edit/:id", middlewares.ApiAuth(token, userService), userHandler.UpdateUser)
	apiV1.DELETE("/delete/:id", middlewares.ApiAuth(token, userService), userHandler.DeleteUser)
	return router
}
