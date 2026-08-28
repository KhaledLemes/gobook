package main

import (
	"gobook/internal/config"
	"gobook/internal/router"
)

func main() {
	config.Load()
	r := router.GerarRouter()

	// FileServer diz que root é a pasta selecionada
	//fs := http.FileServer(http.Dir("./web/templates/"))
	// Handle diz "me passe um caminho e o que ele vai fazer

	r.Run("localhost:8080")
}
