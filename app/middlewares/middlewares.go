package middlewares

import (
	"gop-api/modules/user/service"
	apiresponse "gop-api/package/api-response"
	jwttoken "gop-api/package/jwt-token"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/time/rate"
)

func ApiAuth(jwtAuth jwttoken.JwtToken, userService service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("access-token")

		token, err := jwtAuth.ValidateJwt(header)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid Token",
			})
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid Token",
			})
			return
		}

		user_id := claim["user_id"].(string)

		user, err := userService.GetUserById(user_id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid Token",
			})
			return
		}

		if user.User_id == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid Token",
			})
			return
		}

		c.Set("current", user)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RateLimiter() gin.HandlerFunc {
	limit := rate.NewLimiter(2, 4)
	return func(c *gin.Context) {
		if !limit.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, apiresponse.ResponStatus("Too Many Request", http.StatusTooManyRequests))
			return
		}
		c.Next()
	}
}
