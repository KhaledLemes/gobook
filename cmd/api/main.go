package main

import (
	"gobook/internal/config"
	"gobook/internal/router"
)

func main() {
	config.Load()
	r := router.GerarRouter()
	r.Static("/css", "./web/static/css")
	r.Static("/img", "./web/static/img")
	r.Static("/js", "./web/static/js")

	r.LoadHTMLGlob("web/templates/**/*.html")

	// FileServer diz que root é a pasta selecionada
	//fs := http.FileServer(http.Dir("./web/templates/"))
	// Handle diz "me passe um caminho e o que ele vai fazer

	r.Run("0.0.0.0:8080")
}
