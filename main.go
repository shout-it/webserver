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
         c.JSON(200,gin.H {"Error": err})
      }
      c.JSON(200, gin.H{"status": "Ok"})
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