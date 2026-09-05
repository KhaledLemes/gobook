package router

import (
	controller "gobook/internal/handlers"
	"gobook/internal/middleware"
	"time"

	"github.com/gin-contrib/cors"
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

	// Permite testar na minha máquina pela porta :8081 sem bloqueio de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "http://127.0.0.1:8080", "http://localhost:8081", "http://127.0.0.1:8081"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Lenght", "Content-Type", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	paginas := r.Group("/")
	{
		paginas.GET("/", controller.PaginaConstrucao)
		paginas.GET("/home", controller.PaginaInicial)
		paginas.GET("/login", controller.PaginaLogin)
		paginas.GET("/registro", controller.PaginaRegistro)

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
		protegidos.GET("/me", controller.Me)
		protegidos.GET("/logout", controller.Logout)
	}
	return r
}

func GerarRouter() *gin.Engine {
	r := gin.Default()
	return ConfigRouter(r)
}
