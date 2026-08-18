package models

type Quarto struct {
	Nome       string
	ValorNoite float64

	Disponivel bool

	Propriedade Propriedade
}
