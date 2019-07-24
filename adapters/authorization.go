package adapters

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webserver/helpers"
)

func WithAuth(handler func(c *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		tokenString,err := c.Cookie("token")
		if err != nil {
			c.String(http.StatusUnauthorized,"You are not allowed to access this")
			return
		}
		parsedToken,err := helpers.ParseTokenFromSignedTokenString(tokenString)
		if err != nil {
			c.String(http.StatusUnauthorized,"You are not allowed to access this")
			return
		}
		c.Set("claims",parsedToken.Claims)
		handler(c)
	}
}