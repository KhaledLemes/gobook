package models

import (
	"errors"
	"slices"
	"strings"
)

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

func (p *Propriedade) Prepara() error {
	if err := p.valida(); err != nil {
		return err
	}

	if err := p.formata(); err != nil {
		return err
	}
	return nil
}

func (p *Propriedade) valida() error {
	if p.Nome == "" {
		return errors.New("o nome é mandatório")
	}
	if len(p.Nome) < 6 {
		return errors.New("o nome da propriedade deve ter pelo menos 5 caracteres")
	}

	if p.Descricao == "" {
		return errors.New("a descrição é mandatória")
	}
	if len(p.Descricao) < 30 {
		return errors.New("a descrição da propriedade deve ter pelo menos 30 caracteres")
	}
	if p.Estado == "" {
		return errors.New("o estado é mandatório")
	}

	if ok := slices.Contains(EstadosBrasil, p.Estado); !ok {
		return errors.New("estado inválido")
	}

	if p.Cidade == "" {
		return errors.New("a cidade é mandatória")
	}
	if p.Categoria == "" {
		return errors.New("a categoria é mandatória")
	}
	if ok := slices.Contains(Categorias, p.Categoria); !ok {
		return errors.New("a categoria é é inválida")
	}

	return nil
}

func (p *Propriedade) formata() error {
	p.Nome = strings.TrimSpace(p.Nome)
	p.Descricao = strings.TrimSpace(p.Descricao)
	p.Estado = strings.TrimSpace(p.Estado)
	p.Cidade = strings.TrimSpace(p.Cidade)
	return nil
}
