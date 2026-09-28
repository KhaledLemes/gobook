package models

import (
	"errors"
	"strings"
	"time"
)

type Quarto struct {
	ID int `json:"id"`
	Nome       string    `json:"nome"`
	Descricao  string    `json:"descricao"`
	ValorNoite float64   `json:"valor_noite"`
	Reembolso  time.Time `json:"reembolso"`
	Foto       string    `json:"foto"`

	Disponivel int `json:"disponivel"`

	Propriedade int `json:"propriedad"`

	Dono int `json:"dono"`

}

func (q *Quarto) Validar() error {
	nomeTrim := strings.TrimSpace(q.Nome)
	if nomeTrim == "" {
		return errors.New("o nome do quarto não pode ser vazio")
	}
	if len(nomeTrim) < 3 || len(nomeTrim) > 50 {
		return errors.New("o nome do quarto deve ter entre 3 e 50 caracteres")
	}
	q.Nome = nomeTrim

	if strings.TrimSpace(q.Descricao) == "" {
		return errors.New("a descrição do quarto não pode ser vazia")
	}
	q.Descricao = strings.TrimSpace(q.Descricao)

	if q.ValorNoite < 0 {
		return errors.New("o valor da noite não pode ser menor que 0")
	}

	if q.Disponivel < 0 {
		return errors.New("a quantidade disponível não pode ser menor que 0")
	}

	return nil
}
