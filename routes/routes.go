package routes

import (
	"github.com/gin-gonic/gin"
	. "webserver/adapters"
	. "webserver/constants"
	. "webserver/handlers"
)

type RouteConfig struct {
	Method string
	Path string
	Handler gin.HandlerFunc
}

func GetAllRoutes() []RouteConfig{
	return []RouteConfig {
		{
			Method:HttpPostMethod,
			Path:SignInPathName,
			Handler:SignInHandler,
		},
		{
			Method:HttpPostMethod,
			Path:SignUpPathName,
			Handler:SignUpHandler,
		},
		{
			Method:HttpGetMethod,
			Path:WelcomePathName,
			Handler:WithAuth(WelcomeHandler),
		},
	}
}