package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/api/kk", handler)

	// FileServer diz que root é a pasta selecionada
	fs := http.FileServer(http.Dir("./web/templates/"))
	// Handle diz "me passe um caminho e o que ele vai fazer
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}
