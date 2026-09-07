package middleware

import "errors"

var (
	ErrRoleInvalido      error = errors.New("somente proprietários podem acessar esse recurso")
	ErrRoleInvalidoAdmin error = errors.New("você não tem permissão para acessar esse recurso")
)
