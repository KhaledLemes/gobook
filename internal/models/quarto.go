package models

type Quarto struct {
	Nome       string  `json:"nome"`
	ValorNoite float64 `json:"valor_noite"`

	Disponivel bool `json:"disponivel"`

	Propriedade Propriedade `json:"propriedad"`
}
