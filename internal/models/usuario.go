package models

import (
	"errors"
	"strings"
	"time"

	"github.com/KhaledLemes/gobook/internal/security"
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
	ID int32

	Nome       string
	NomeMeio   string
	NomeUltimo string
	Nascimento time.Time

	Email string
	Senha string

	Role Role

	NumReservas int32
	Reservas    []Reserva

	CriadoEm time.Time
}

// Prepara valida e salva os dados do usuário em uma struct chamada Usuario
// Recebe um valor booleano SeRegistrando, do qual muda seu comportamento caso esteja sendo invocado na hora do login. Isso permite que a gente usa a mesma função tanto para login quanto para registro.
func (u *Usuario) Prepara(SeRegistrando bool, role Role) error {

	if err := u.validar(SeRegistrando); err != nil {
		return err
	}

	if err := u.formatar(SeRegistrando, role); err != nil {
		return err
	}

	return nil
}

func (u *Usuario) validar(SeRegistrando bool) error {
	if u.Nome == "" {
		return errors.New("O nome é mandatório")
	}
	if u.NomeMeio == "" {
		return errors.New("O nome é mandatório")
	}
	if u.NomeUltimo == "" {
		return errors.New("O último nome é mandatório")
	}
	if u.Email == "" {
		return errors.New("O campo de e-mail é mandatório")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("Formato do e-mail inváldio")
	}

	if SeRegistrando && u.Senha == "" {
		return errors.New("A senha é um campo obrigatório")
	}

	return nil
}

func (u *Usuario) formatar(SeRegistrando bool, role Role) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.NomeMeio = strings.TrimSpace(u.NomeMeio)
	u.NomeUltimo = strings.TrimSpace(u.NomeUltimo)
	u.Role = role

	if SeRegistrando {
		senhaHash, err := security.Hash(u.Senha)
		if err != nil {
			return err
		}
		u.Senha = string(senhaHash)
	}

	return nil
}
