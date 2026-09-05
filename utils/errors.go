package utils

import "strings"

// VerificaErro retorna true caso comp exista em err. Criei para evitar repetição de lógica com poluição visual
func VerificaErro(err error, comp string) bool {
	return strings.Contains(err.Error(), comp)
}
