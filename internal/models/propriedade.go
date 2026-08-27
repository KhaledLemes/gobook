package models

import (
	"errors"
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
	if p.Descricao == "" {
		return errors.New("a descrição é mandatória")
	}
	if p.Estado == "" {
		return errors.New("o estado é mandatório")
	}
	if p.Cidade == "" {
		return errors.New("a cidade é mandatória")
	}
	if p.Categoria == "" {
		return errors.New("o categoria é mandatória")
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
