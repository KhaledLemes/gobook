package models

import (
	"errors"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Categoria string

const (
	Hotel       Categoria = "hotel"
	Casa        Categoria = "casa"
	Apartamento Categoria = "apartamento"
	Chale       Categoria = "chalé"
	Pousada     Categoria = "pousada"
)

const rex string = `^[\p{L}\s]+$`

type Propriedade struct {
	ID            int64  `json:"id"`
	Nome          string `json:"nome"`
	Endereco      string `json:"endereco"`
	Numero        string `json:"numero"`
	Descricao     string `json:"descricao"`
	Estado        string `json:"estado"`
	Cidade        string `json:"cidade"`
	FotoPrincipal string `json:"foto"`

	Avaliacao float64 `json:"avaliacao"`

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
	if len(p.Nome) > 200 {
		return errors.New("o nome da propriedade é muito grande")
	}

	if p.Descricao == "" {
		return errors.New("a descrição é mandatória")
	}
	if len(p.Descricao) < 30 {
		return errors.New("a descrição da propriedade deve ter pelo menos 30 caracteres")
	}
	if len(p.Descricao) > 1000 {
		return errors.New("a descrição da propriedade deve ter no máximo 2000 caracteres")
	}

	if p.Categoria == "" {
		return errors.New("a categoria é mandatória")
	}
	catStr := strings.ToLower(strings.TrimSpace(string(p.Categoria)))
	p.Categoria = Categoria(catStr)
	if ok := slices.Contains(Categorias, p.Categoria); !ok {
		return errors.New("a categoria é é inválida")
	}

	if p.Endereco == "" {
		return errors.New("o endereco é mandatório")
	}
	if len(p.Endereco) < 6 {
		return errors.New("o campo com o nome da rua deve conter no mínimo 6 caracteres")
	}
	if len(p.Endereco) > 264 {
		return errors.New("o campo com o nome da rua deve conter no máximo 264 caracteres")
	}
	if ok, err := regexp.Match(rex, []byte(p.Endereco)); !ok || err != nil {
		return errors.New("endereço inválido")
	}

	if p.Numero == "" {
		return errors.New("número da propriedade é um campo obrigatório")
	}
	if len(p.Numero) > 5 {
		return errors.New("número da propriedade é inválido")
	}
	if _, err := strconv.Atoi(p.Numero); err != nil {
		return errors.New("número da propriedade é inválido")
	}

	if p.Cidade == "" {
		return errors.New("a cidade é mandatória")
	}
	if len(p.Cidade) < 3 {
		return errors.New("a cidade deve conter no mínimo 3 caracteres")
	}
	if len(p.Cidade) > 50 {
		return errors.New("a cidade deve conter no máximo 50 caracteres")
	}

	if p.Estado == "" {
		return errors.New("o estado é mandatório")
	}
	p.Estado = strings.ToUpper(strings.TrimSpace(p.Estado))
	if ok := slices.Contains(UFsBrasil, p.Estado); !ok {
		return errors.New("estado inválido")
	}

	return nil
}

func (p *Propriedade) formata() error {
	p.Nome = strings.TrimSpace(p.Nome)
	p.Endereco = strings.TrimSpace(p.Endereco)
	p.Numero = strings.TrimSpace(p.Numero)
	p.Descricao = strings.TrimSpace(p.Descricao)
	p.Estado = strings.TrimSpace(p.Estado)
	p.Cidade = strings.TrimSpace(p.Cidade)

	p.Nome = strings.ToLower(p.Nome)
	p.Endereco = strings.ToLower(p.Endereco)
	p.Cidade = strings.ToLower(p.Cidade)

	p.Nome = primeiraLetraToUpper(p.Nome)
	p.Endereco = primeiraLetraToUpper(p.Endereco)
	p.Cidade = primeiraLetraToUpper(p.Cidade)

	return nil
}
