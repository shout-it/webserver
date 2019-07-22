package main
import (
   "net/http"
   "github.com/gin-gonic/gin"
)
func main() {
   router := SetupRouter()
   router.Run()
}
func SetupRouter() *gin.Engine {
   router := gin.Default()
   router.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{
          "hello": "world",
      })
   })
   
   return router
}