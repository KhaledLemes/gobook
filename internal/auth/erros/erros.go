package erros

import "errors"

var (
	ErrSessaoExpirada error = errors.New("sessão expirada, por favor, faça login novamente")
	ErrTknInvalido    error = errors.New("token inválido")
)
