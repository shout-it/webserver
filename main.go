package main

import (
   "github.com/gin-gonic/gin"
   "webserver/handlers"
)
var jwtKey = []byte("aa9060d3-b56d-4c52-b81e-2edb06ed6697")

func main() {
   router := SetupRouter()
   router.Run()
}
func SetupRouter() *gin.Engine {
   router := gin.Default()
   router.POST("/signup",handlers.SignUpHandler)
   router.POST("/signin", handlers.SignInHandler)
   return router
}


