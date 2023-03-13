package routes

import (
	"gop-api/app/middlewares"
	"gop-api/modules/user/controller"
	"gop-api/modules/user/repository"
	"gop-api/modules/user/service"
	htmlrender "gop-api/package/html-render"
	jwttoken "gop-api/package/jwt-token"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	token := jwttoken.NewJwtToken()
	router.HTMLRender = htmlrender.Render("./public/templates")

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := controller.NewUserHandler(userService, token)
	pagesView := controller.NewPagesView(userService)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Endpoint Not Found",
		})
	})

	router.GET("/", pagesView.Indexview)
	router.GET("/others", pagesView.Others)

	router.POST("/login", userHandler.Login)
	router.POST("/register", userHandler.AddUser)

	apiV1 := router.Group("api/v1/user", middlewares.ApiAuth(token, userService), middlewares.RateLimiter())
	apiV1.GET("/", userHandler.User)
	apiV1.GET("/:id", userHandler.DetailUser)
	apiV1.PUT("/:id", userHandler.UpdateUser)
	apiV1.DELETE("/:id", userHandler.DeleteUser)
	return router
}
