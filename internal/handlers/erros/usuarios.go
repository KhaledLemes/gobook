package erros

import "errors"

var (
	ErrCredenciaisInvalidas error = errors.New("credenciais inválidas. Reveja os dados e tente novamente")
	ErrDataNascInvalida     error = errors.New("data de nascimento inválida")
)
