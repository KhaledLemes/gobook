package main

import (
	"gobook/internal/config"
	"gobook/internal/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

func a(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func main() {
	config.Load()
	r := router.GerarRouter()

	// FileServer diz que root é a pasta selecionada
	//fs := http.FileServer(http.Dir("./web/templates/"))
	// Handle diz "me passe um caminho e o que ele vai fazer

	r.LoadHTMLFiles("./web/templates/index.html")
	r.StaticFile("/css", "./web/templates/css/style.css")
	r.GET("/", a)
	//http.Handle("/", fs)
	//http.ListenAndServe(":8080", r)
	r.Run("localhost:8080")
}
