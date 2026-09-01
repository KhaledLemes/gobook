package models

import (
	"errors"
	"strings"
	"time"

	"gobook/internal/security"

	"github.com/badoux/checkmail"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleOwner Role = "owner"
	RoleGuest Role = "guest"
	Forbidden Role = "forbidden"
)

type Usuario struct {
	ID int `json:"id"`

	Nome       string    `json:"nome"`
	NomeMeio   string    `json:"nome_meio"`
	NomeUltimo string    `json:"nome_ultimo"`
	Nascimento time.Time `json:"nascimento"`
	Tel        string    `json:"tel"`

	Credito string `json:"credito"`

	Email string `json:"email"`
	Senha string `json:"senha"`

	Role Role `json:"role"`

	NumReservas int32     `json:"num_reservas"`
	Reservas    []Reserva `json:"reservas"`

	// Esse é definido pelo banco de dados
	CriadoEm time.Time `json:"criado_em"`
}

// Prepara valida e salva os dados do usuário em uma struct chamada Usuario
// Recebe um valor booleano SeRegistrando, do qual muda seu comportamento caso esteja sendo invocado na hora do login. Isso permite que a gente usa a mesma função tanto para login quanto para registro.
func (u *Usuario) Prepara(SeRegistrando bool) error {

	if err := u.validar(SeRegistrando); err != nil {
		return err
	}

	if err := u.formatar(SeRegistrando); err != nil {
		return err
	}

	return nil
}

func (u *Usuario) validar(SeRegistrando bool) error {
	if SeRegistrando && u.Nome == "" {
		return errors.New("O nome é mandatório")
	}
	if SeRegistrando && u.NomeUltimo == "" {
		return errors.New("o último nome é mandatório")
	}
	if u.Email == "" {
		return errors.New("o campo de e-mail é mandatório")
	}

	if SeRegistrando && u.Role != "guest" && u.Role != "owner" {
		return errors.New("tipo de usuário inválido")
	}

	if err := u.validaNascimento(); SeRegistrando && err != nil {
		return err
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("formato do e-mail inválido")
	}

	if u.Senha == "" {
		return errors.New("A senha é um campo obrigatório")
	}

	return nil
}

func (u *Usuario) formatar(SeRegistrando bool) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.NomeMeio = strings.TrimSpace(u.NomeMeio)
	u.NomeUltimo = strings.TrimSpace(u.NomeUltimo)

	if SeRegistrando {
		senhaHash, err := security.Hash(u.Senha)
		if err != nil {
			return err
		}
		u.Senha = string(senhaHash)
	}

	return nil
}

func (u *Usuario) validaNascimento() error {
	if u.Nascimento.After(time.Now()) {
		return errors.New("nascimento inválido")
	}
	if (time.Now().Year() - u.Nascimento.Year()) < 18 {
		return errors.New("novo usuário não pode ser menor de 18 anos")
	}
	return nil
}
