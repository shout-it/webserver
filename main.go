package main

import (
   "github.com/gin-gonic/gin"
   "strconv"
   "webserver/config"
   "webserver/routes"
)

func main() {
   config := config.GetConfig()
   routesConfig := routes.GetAllRoutes()
   router := gin.Default()
   for _,route := range routesConfig {
      router.Handle(route.Method,route.Path,route.Handler)
   }
   router.Run(":" + strconv.Itoa(config.Port))
}
