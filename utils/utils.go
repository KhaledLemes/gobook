package utils

import (
	"strings"
	"unicode/utf8"
)

func PrimeiraLetraToUpper(s string) string {
	primeira, size := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(primeira)) + s[size:]
}
