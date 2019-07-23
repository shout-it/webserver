package main

import (
   "github.com/gin-gonic/gin"
   "webserver/dao"
   "webserver/models"
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
      dao.InsertOneValue(user)
      c.JSON(200, gin.H{"status": "Ok"})
   })
   return router
}