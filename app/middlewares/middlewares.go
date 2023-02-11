package middlewares

import (
	jwttoken "gop-api/app/jwt-token"
	"gop-api/modules/user/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

		if user.User_uuid == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "Invalid Token",
			})
			return
		}

		c.Set("current", user)
	}
}
