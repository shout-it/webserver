package routes

import (
	"github.com/gin-gonic/gin"
	. "webserver/adapters"
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
			Method:"POST",
			Path:"/signin",
			Handler:SignInHandler,
		},
		{
			Method:"POST",
			Path:"/signup",
			Handler:SignUpHandler,
		},
		{
			Method:"POST",
			Path:"/welcome",
			Handler:WithAuth(WelcomeHandler),
		},
	}
}