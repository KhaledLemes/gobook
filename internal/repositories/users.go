package repositories

import (
	"database/sql"
	models "internal/models"
)

type Usuarios struct {
	db *sql.DB
}

func NewUsuariosRepo(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (r Usuarios) CriarNovoUsuario(usuario models.Usuario) {

}
