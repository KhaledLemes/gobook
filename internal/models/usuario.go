package models

import (
	"errors"
	"regexp"
	"slices"
	"strings"
	"time"
	"unicode/utf8"

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

const rexNome string = `^[\p{L}\s]+$`

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
	if u.Email == "" {
		return errors.New("o e-mail é mandatório")
	}
	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return errors.New("formato do e-mail inválido")
	}

	if u.Senha == "" {
		return errors.New("A senha é um campo obrigatório")
	}

	if SeRegistrando {
		if u.Nome == "" {
			return errors.New("o nome é mandatório")
		}
		if ok, err := regexp.Match(rexNome, []byte(u.Nome)); !ok || err != nil {
			return errors.New("revise o nome e tente novamente")
		}

		if ok, err := regexp.Match(rexNome, []byte(u.NomeMeio)); !ok || err != nil {
			if u.NomeMeio != "" {
				return errors.New("revise o nome do meio e tente novamente")
			}
		}

		if u.NomeUltimo == "" {
			return errors.New("o último nome é mandatório")
		}
		if ok, err := regexp.Match(rexNome, []byte(u.NomeUltimo)); !ok || err != nil {
			return errors.New("revise o último nome e tente novamente")
		}

		if u.Role != RoleGuest && u.Role != RoleOwner {
			return errors.New("tipo de usuário inválido")
		}

		if len(u.Senha) < 8 {
			return errors.New("a senha é muito curta")
		}
		if err := validaSenha([]byte(u.Senha)); err != nil {
			return err
		}

		if err := u.validaNascimento(); err != nil {
			return err
		}
	}
	return nil
}

func (u *Usuario) formatar(SeRegistrando bool) error {
	u.Nome = strings.TrimSpace(u.Nome)
	u.NomeMeio = strings.TrimSpace(u.NomeMeio)
	u.NomeUltimo = strings.TrimSpace(u.NomeUltimo)
	u.Email = strings.TrimSpace(u.Email)

	u.Nome = strings.ToLower(u.Nome)
	u.NomeMeio = strings.ToLower(u.NomeMeio)
	u.NomeUltimo = strings.ToLower(u.NomeUltimo)
	u.Email = strings.ToLower(u.Email)

	u.Nome = primeiraLetraToUpper(u.Nome)
	u.NomeMeio = primeiraLetraToUpper(u.NomeMeio)
	u.NomeUltimo = primeiraLetraToUpper(u.NomeUltimo)


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
	if u.Nascimento.IsZero() {
		return errors.New("a data de nascimento é um campo obrigatório")
	}
	if u.Nascimento.After(time.Now()) {
		return errors.New("nascimento inválido")
	}
	if (time.Now().Year() - u.Nascimento.Year()) < 18 {
		return errors.New("novo usuário não pode ser menor de 18 anos")
	}
	if (time.Now().Year() - u.Nascimento.Year()) > 110 {
		return errors.New("a idade é inválida")
	}
	return nil
}

func validaSenha(ps []byte) error {
	var number, upperCase, lowerCase, specialDigit bool

	for _, c := range ps {
		switch true {
		case c >= 48 && c <= 57:
			number = true
			break
		case c >= 65 && c <= 90:
			upperCase = true
			break
		case c >= 97 && c <= 122:
			lowerCase = true
			break
		case especial(c):
			specialDigit = true
			break
		}
	}

	if number && upperCase && lowerCase && specialDigit {
		return nil
	}
	return errors.New("a senha deve conter ao menos um número, letra minúscula, letra maiúscula e dígito especial (!, #, $, %, &, *, +, -, ., ?, @)")
}

// verifica se char é um dos seguintes:
// ! # $ % & * + - . ? @
func especial(n byte) bool {
	var lista = []byte{33, 35, 36, 37, 38, 42, 43, 45, 46, 63, 64}
	return slices.Contains(lista, n)
}

func primeiraLetraToUpper(s string) string {
	primeira, size := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(primeira)) + s[size:]
}
