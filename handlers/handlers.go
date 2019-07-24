package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	. "webserver/constants"
	"webserver/dao"
	"webserver/helpers"
	"webserver/models"
)

func WelcomeHandler(c *gin.Context) {
	value,_ := c.Get("claims")
	c.JSON(http.StatusOK,gin.H{"Claims":value})
}

func SignUpHandler(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	pwd := []byte(user.Password)
	hashedPassword := helpers.HashAndSalt(pwd)
	user.Password = hashedPassword
	err := dao.InsertOneValue(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError,err)
		return
	}
	c.JSON(http.StatusOK, "Created the resource")
	return
}

func SignInHandler(c *gin.Context) {
	var credentials models.Credentials
	c.BindJSON(&credentials)
	user,err := dao.FindBy(credentials.Email)
	if err != nil {
		c.JSON(404,"Could not find Email")
		return
	}
	matched := helpers.ComparePasswords(user.Password,[]byte(credentials.Password))
	if matched {
		expirationTime := time.Now().Add(5*time.Minute);
		tokenString,err := helpers.CreateSignedTokenStringFor(expirationTime.Unix(),credentials)
		if err != nil {
			c.JSON(http.StatusInternalServerError,"Failed to generate signed token")
			return
		}
		c.SetCookie(AuthTokenName,tokenString,int(expirationTime.Unix()),"","",false,false)
		return
	}
	c.JSON(http.StatusUnauthorized,"Invalid Credentials")
	return
}

