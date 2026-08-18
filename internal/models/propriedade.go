package models

type Categoria string

const (
	Hotel       Categoria = "Hotel"
	Pousada     Categoria = "Casa"
	Apartamento Categoria = "Apartamento"
	Chale       Categoria = "Chalé"
)

type Propriedade struct {
	ID        int64
	Nome      string
	Descricao string
	Estado    string
	Cidade    string

	PetFriendly bool

	Categoria Categoria
	Dono      Usuario

	Quartos []Quarto
}
