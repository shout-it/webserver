package main

import (
   "github.com/gin-gonic/gin"
   "log"
   "webserver/dao"
   "webserver/models"
   "golang.org/x/crypto/bcrypt"
)
func main() {
   router := SetupRouter()
   router.Run()
}
func SetupRouter() *gin.Engine {
   router := gin.Default()
   router.POST("/signup", func(c *gin.Context) {
      var user models.User
      c.BindJSON(&user)
      pwd := []byte(user.Password)
      hashedPassword := hashAndSalt(pwd)
      user.Password = hashedPassword
      err := dao.InsertOneValue(user)
      if err != nil {
         c.JSON(500,gin.H {"Error": err})
      }
      c.JSON(200, gin.H{"status": "Ok"})
   })
   router.POST("/login", func(c *gin.Context) {
      var credentials models.Credentials
      c.BindJSON(&credentials)
      user,err := dao.FindBy(credentials.Email)
      if err != nil {
         c.JSON(500,gin.H{"Error": "Could not find Email"})
      } else {
         matched := comparePasswords(user.Password,[]byte(credentials.Password))
         if matched {
            c.JSON(200,gin.H{"result":"Authentication Successfull"})
         } else {
            c.JSON(403,gin.H{"result":"Failed to Authenticate"})
         }
      }

   })
   return router
}


func hashAndSalt(pwd []byte) string {
   hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
   if err != nil {
      log.Println(err)
   }
   return string(hash)
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
   byteHash := []byte(hashedPwd)
   err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
   if err != nil {
      log.Println(err)
      return false
   }
   return true
}