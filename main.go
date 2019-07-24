package main

import (
   "github.com/gin-gonic/gin"
   "webserver/handlers"
)

func main() {
   router := SetupRouter()
   router.Run()
}
func SetupRouter() *gin.Engine {
   router := gin.Default()
   router.POST("/signup",handlers.SignUpHandler)
   router.POST("/signin", handlers.SignInHandler)
   router.POST("/story",handlers.WithAuth(handlers.StoryHadler) )
   return router
}


