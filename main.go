package main

import (
   "github.com/gin-gonic/gin"
   "strconv"
   "webserver/adapters"
   "webserver/config"
   "webserver/handlers"
)

func main() {
   config := config.GetConfig()
   router := SetupRouter()
   router.Run(":" + strconv.Itoa(config.Port))
}
func SetupRouter() *gin.Engine {
   router := gin.Default()
   router.POST("/signup",handlers.SignUpHandler)
   router.POST("/signin", handlers.SignInHandler)
   router.POST("/welcome", adapters.WithAuth(handlers.WelcomeHandler) )
   return router
}


