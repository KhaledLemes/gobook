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

	paginasPublicas := r.Group("/")
	{
		paginasPublicas.GET("/", controller.PaginaConstrucao)
		paginasPublicas.GET("/home", controller.PaginaInicial)
		paginasPublicas.GET("/login", controller.PaginaLogin)
		paginasPublicas.GET("/registro", controller.PaginaRegistro)

		paginasPublicas.GET("/unauthorized", controller.Unauthorized)
	}

	paginasProprietarios := r.Group("/propriedades")
	paginasProprietarios.Use(middleware.AutenticaProprietario())
	{
		paginasProprietarios.GET("/minhas", controller.PaginaOwner)
		paginasProprietarios.GET("/nova", controller.PaginaCriarPropriedade)
		paginasProprietarios.GET("/editar", controller.PaginaOwner)
		paginasProprietarios.GET("/excluir", controller.PaginaOwner)

	}
	paginasAdmin := r.Group("/admin")
	paginasAdmin.Use(middleware.AutenticaAdmin())
	{
		paginasAdmin.GET("painel", controller.PainelAdmin)
	}

	rotasPublicas := r.Group("/api/v1")
	{
		rotasPublicas.POST("/login", controller.Login)

		rotasPublicas.POST("/usuarios", controller.CriaUsuario)

		rotasPublicas.GET("/propriedades", controller.MostraTodasPropriedades)
		rotasPublicas.GET("/propriedades/iniciais", controller.BuscarDezPrimeirasPropriedadesAleatorio)
		rotasPublicas.GET("/propriedades/id/:id", controller.BuscaPropriedadePorID)
		rotasPublicas.GET("/propriedades/:nome", controller.BuscaPropriedadePorNome)

	}

	rotasProtegidas := r.Group("/api/v1")
	rotasProtegidas.Use(middleware.Autentica())
	{
		rotasProtegidas.POST("/propriedades", controller.CriarPropriedade)
		rotasProtegidas.PUT("/propriedades/:id", controller.EditarPropriedade)
		rotasProtegidas.DELETE("/propriedades/:id", controller.DeletaPropriedadePorID)
		rotasProtegidas.GET("/propriedades/minhas", controller.BuscarTodasPorDono)

		rotasProtegidas.GET("/me", controller.Me)
		rotasProtegidas.GET("/logout", controller.Logout)

	}

	return r
}

func GerarRouter() *gin.Engine {
	r := gin.Default()
	return ConfigRouter(r)
}
