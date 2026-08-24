package router

import (
	"gobook/internal/middleware"

	"github.com/gin-gonic/gin"
)

type Route struct {
	URI        string
	Metodo     string
	Func       gin.HandlerFunc
	RequerAuth bool
}

func ConfigRouter(r *gin.Engine) *gin.Engine {
	routes := usersRoutes
	routes = append(routes, loginRoute)

	for _, r := range propriedadesRouter {
		routes = append(routes, r)
	}

	for _, route := range routes {
		if route.RequerAuth {
			r.Handle(route.Metodo, route.URI,
				middleware.Logger(
					middleware.Autentica(route.Func)))
		} else {
			r.Handle(route.Metodo, route.URI,
				middleware.Logger(route.Func))
		}
	}
	return r
}

func GerarRouter() *gin.Engine {
	r := gin.Default()
	return ConfigRouter(r)
}
