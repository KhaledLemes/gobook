package utils

import (
	"errors"
	"net/http"
	"strings"
	"unicode/utf8"
)

func PrimeiraLetraToUpper(s string) string {
	primeira, size := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(primeira)) + s[size:]
}

func ValidarFotoStream(fileBytes []byte) error {
	if len(fileBytes) == 0 {
		return errors.New("o arquivo de foto não pode ser vazio")
	}

	// Lê os primeiros 512 bytes para detectar o Content-Type real do arquivo
	mimeType := http.DetectContentType(fileBytes)

	// Garante que o tipo detectado comece com "image/" (ex: image/jpeg, image/png, image/webp)
	if !strings.HasPrefix(mimeType, "image/") {
		return errors.New("o arquivo enviado não é uma imagem válida")
	}

	return nil
}
