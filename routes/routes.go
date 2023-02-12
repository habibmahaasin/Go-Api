package routes

import (
	"gop-api/app/middlewares"
	"gop-api/modules/user/controller"
	"gop-api/modules/user/repository"
	"gop-api/modules/user/service"
	jwttoken "gop-api/package/jwt-token"
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
			"message": "Endpoint Not Found",
		})
	})

	apiV1 := router.Group("api/v1/user")
	apiV1.GET("/", middlewares.ApiAuth(token, userService), userHandler.User)
	apiV1.POST("/login", userHandler.Login)
	apiV1.POST("/register", userHandler.AddUser)
	apiV1.GET("/:id", middlewares.ApiAuth(token, userService), userHandler.DetailUser)
	apiV1.PUT("/:id", middlewares.ApiAuth(token, userService), userHandler.UpdateUser)
	apiV1.DELETE("/:id", middlewares.ApiAuth(token, userService), userHandler.DeleteUser)
	return router
}
