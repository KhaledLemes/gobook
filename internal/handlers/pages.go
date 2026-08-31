package handlers

import "github.com/gin-gonic/gin"

func PaginaConstrucao(c *gin.Context) {
	data := gin.H{
		"title": "Gobook - Construindo!",
	}
	c.HTML(200, "construcao.html", data)
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
