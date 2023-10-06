package middleware

import (
	"net/http"

	"api/utils/token"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			// if expired cookie present, remove it
			c.SetCookie("Autorization", "", -1, "/", "", false, true)
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
