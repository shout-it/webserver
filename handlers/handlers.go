package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
	"webserver/dao"
	"webserver/helpers"
	"webserver/models"
)

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
			claims := &models.Claims {
				Email: credentials.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt:expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodES256,claims)
			tokenString,err := token.SigningString()
			if err != nil {
				c.JSON(500,gin.H{"Error":err})
			}
			c.SetCookie("token",tokenString,int(expirationTime.Unix()),"","",false,false)
		} else {
			c.JSON(403,gin.H{"result":"Failed to Authenticate"})
		}
	}
}

