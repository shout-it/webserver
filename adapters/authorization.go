package adapters

import (
	"github.com/gin-gonic/gin"
	"net/http"
	. "webserver/constants"
	"webserver/helpers"
)

func WithAuth(handler func(c *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		tokenString,err := c.Cookie(AuthTokenName)
		if err != nil {
			c.String(http.StatusUnauthorized,"You are not allowed to access this")
			return
		}
		parsedToken,err := helpers.ParseTokenFromSignedTokenString(tokenString)
		if err != nil {
			c.String(http.StatusUnauthorized,"You are not allowed to access this")
			return
		}
		c.Set(Claims,parsedToken.Claims)
		handler(c)
	}
}