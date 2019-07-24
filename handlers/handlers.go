package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"webserver/dao"
	"webserver/helpers"
	"webserver/models"
)

func WithAuth(adapter func(c *gin.Context)) func(c *gin.Context){
	return func(c *gin.Context) {
		tokenString,err := c.Cookie("token")
		if err != nil {
			c.String(http.StatusUnauthorized,"You are not allowed to access this")
			return
		}
		parsedToken,err := helpers.ParseTokenFromSignedTokenString(tokenString)
		if err != nil {
			log.Print(err,"hello")
			c.String(http.StatusUnauthorized,"You are not allowed to access this")
			return
		}
		c.Set("userInfo",parsedToken)
		adapter(c)
	}
}

func StoryHadler(c *gin.Context) {
	c.JSON(200,"Stories")
	return
}

func SignUpHandler(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	pwd := []byte(user.Password)
	hashedPassword := helpers.HashAndSalt(pwd)
	user.Password = hashedPassword
	err := dao.InsertOneValue(user)
	if err != nil {
		c.JSON(500,gin.H {"Error": err})
	}
	c.JSON(200, gin.H{"status": "Ok"})
}

func SignInHandler(c *gin.Context) {
	var credentials models.Credentials
	c.BindJSON(&credentials)
	user,err := dao.FindBy(credentials.Email)
	if err != nil {
		c.JSON(500,gin.H{"Error": "Could not find Email"})
	} else {
		matched := helpers.ComparePasswords(user.Password,[]byte(credentials.Password))
		if matched {
			expirationTime := time.Now().Add(5*time.Minute);
			tokenString,err := helpers.CreateSignedTokenStringFor(expirationTime.Unix(),credentials)
			if err != nil {
				c.JSON(500,gin.H{"Error": err})
			}
			c.SetCookie("token",tokenString,int(expirationTime.Unix()),"","",false,false)
		} else {
			c.JSON(403,gin.H{"result":"Failed to Authenticate"})
		}
	}
}

