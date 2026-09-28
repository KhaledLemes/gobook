package handlers

import (
	"fmt"
	"gobook/internal/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PaginaConstrucao(c *gin.Context) {
	data := gin.H{
		"title": "Gobook - Construindo!",
	}
	c.HTML(200, "construcao.html", data)
}

func NotFound(c *gin.Context) {
	data := gin.H{
		"title": "Gobook - Página não encontrada!",
	}
	c.HTML(404, "notFound.html", data)
}

func PaginaLogin(c *gin.Context) {
	data := gin.H{
		"title": "Gobook - Cadastro e Login",
	}
	c.HTML(200, "login.html", data)
}

func PaginaInicial(c *gin.Context) {
	data := gin.H{
		"title": "Gobook - Home",
	}
	c.HTML(200, "home.html", data)
}

func PaginaRegistro(c *gin.Context) {
	data := gin.H{
		"title": "Gobook - Registro",
	}
	c.HTML(200, "registro.html", data)
}

func PaginaReserva(c *gin.Context) {
	propertyName := c.Param("propriedade")
	title := fmt.Sprintf("Gobook - %s!", propertyName)
	data := gin.H{
		"title": title,
	}
	c.HTML(200, "propriedade.html", data)
}

func PaginaOwner(c *gin.Context) {
	dataUnauth := gin.H{
		"title": "Gobook - Invasor!",
	}
	dataAuth := gin.H{
		"title": "Gobook - Painel do proprietário",
	}

	if _, err := c.Cookie("auth"); err != nil {
		c.HTML(http.StatusUnauthorized, "unauthorized.html", dataUnauth)
		return
	}

	role, err := auth.PegarRoleUsuario(c)
	if err != nil || role != "owner" && role != "admin" {
		c.HTML(http.StatusUnauthorized, "unauthorized.html", dataUnauth)
		return
	}

	c.HTML(http.StatusOK, "owner-panel.html", dataAuth)
}

func PropriedadeOwner(c *gin.Context) {
	dataUnauth := gin.H{
		"title": "Gobook - Invasor!",
	}
	dataAuth := gin.H{
		"title": "Gobook - Sua propriedade",
	}

	if _, err := c.Cookie("auth"); err != nil {
		c.HTML(http.StatusUnauthorized, "unauthorized.html", dataUnauth)
		return
	}

	role, err := auth.PegarRoleUsuario(c)
	if err != nil || role != "owner" && role != "admin" {
		c.HTML(http.StatusUnauthorized, "unauthorized.html", dataUnauth)
		return
	}

	c.HTML(http.StatusOK, "owner-propriedade.html", dataAuth)
}

func NovoQuarto(c *gin.Context) {
	dataUnauth := gin.H{
		"title": "Gobook - Novo quarto!",
	}
	dataAuth := gin.H{
		"title": "Gobook - Sua propriedade",
	}

	if _, err := c.Cookie("auth"); err != nil {
		c.HTML(http.StatusUnauthorized, "unauthorized.html", dataUnauth)
		return
	}

	role, err := auth.PegarRoleUsuario(c)
	if err != nil || role != "owner" && role != "admin" {
		c.HTML(http.StatusUnauthorized, "unauthorized.html", dataUnauth)
		return
	}

	c.HTML(http.StatusOK, "adicionar-quarto.html", dataAuth)
}

func Unauthorized(c *gin.Context) {
	data := gin.H{
		"title": "Gobook - Invasor!",
	}
	c.HTML(http.StatusUnauthorized, "unauthorized.html", data)
	return
}

func PainelAdmin(c *gin.Context) {
	dataUnauth := gin.H{
		"title": "Gobook - Invasor!",
	}
	dataAuth := gin.H{
		"title": "Gobook - Painel do admin",
	}

	role, err := auth.PegarRoleUsuario(c)
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
		c.HTML(http.StatusUnauthorized, "unauthorized.html", dataUnauth)
		return
	}

	if role != "admin" {
		c.HTML(http.StatusForbidden, "unauthorized.html", dataUnauth)
		return
	}

	c.HTML(http.StatusOK, "admin.html", dataAuth)
}

func PaginaCriarPropriedade(c *gin.Context) {
	data := gin.H{
		"title": "Gobook - Criar propriedade!",
	}
	c.HTML(http.StatusOK, "propriedade-nova.html", data)
	return
}
