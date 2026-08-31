package router

import (
	controller "gobook/internal/handlers"
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
	r.Use(middleware.Logger())

	paginas := r.Group("/")
	{
		paginas.GET("/", controller.PaginaConstrucao)
		paginas.GET("/home", controller.PaginaInicial)
		paginas.GET("/login", controller.PaginaLogin)
	}
	publicos := r.Group("/api/v1")
	{
		publicos.POST("/login", controller.Login)

		publicos.POST("/usuarios", controller.CriaUsuario)

		publicos.GET("/propriedades", controller.MostraTodasPropriedades)
		publicos.GET("/propriedades/id/:id", controller.BuscaPropriedadePorID)
		publicos.GET("/propriedades/:nome", controller.BuscaPropriedadePorNome)
	}

	protegidos := r.Group("/api/v1")
	protegidos.Use(middleware.Autentica())
	{
		protegidos.POST("/propriedades", controller.CriarPropriedade)
		protegidos.PUT("/propriedades/:id", controller.EditarPropriedade)
		protegidos.DELETE("/propriedades/:id", controller.DeletaPropriedadePorID)
	}
	return r
}

func GerarRouter() *gin.Engine {
	r := gin.Default()
	return ConfigRouter(r)
}
