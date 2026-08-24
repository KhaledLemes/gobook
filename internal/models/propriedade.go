package models

type Categoria string

const (
	Hotel       Categoria = "hotel"
	Pousada     Categoria = "casa"
	Apartamento Categoria = "apartamento"
	Chale       Categoria = "chalé"
)

type Propriedade struct {
	ID        int64  `json:"id"`
	Nome      string `json:"nome"`
	Descricao string `json:"descricao"`
	Estado    string `json:"estado"`
	Cidade    string `json:"cidade"`

	PetFriendly bool `json:"pet_friendly"`

	Categoria Categoria `json:"categoria"`
	Dono      Usuario   `json:"dono"`

	Quartos []Quarto `json:"quartos"`
}
