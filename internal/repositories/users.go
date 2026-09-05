package repositories

import (
	"database/sql"
	"errors"
	_ "gobook/internal/database"
	models "gobook/internal/models"
)

type RepoUsuarios struct {
	db *sql.DB
}

func NewUsuariosRepo(db *sql.DB) *RepoUsuarios {
	return &RepoUsuarios{db}
}

func (r RepoUsuarios) CriarNovoUsuario(usuario models.Usuario) (int, error) {
	stmt, err := r.db.Prepare(
		"INSERT INTO usuarios (nome, nome_do_meio, ultimo_nome, nascimento, email, senha, role, numero_reservas) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
	)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(usuario.Nome, usuario.NomeMeio, usuario.NomeUltimo, usuario.Nascimento, usuario.Email, usuario.Senha, usuario.Role, usuario.NumReservas)
	if err != nil {
		return -1, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(lastInsertID), nil

}

// ProcurarPorEmail retorna ID e senha do usuário procurado para que posteriormente seja comparado com a senha do DB.
// O ID servirá posteriormente para token jwt
func (r RepoUsuarios) ProcurarPorEmail(email string) (models.Usuario, error) {
	rows, err := r.db.Query("SELECT id, senha, role, nome FROM usuarios WHERE email = ?", email)
	if err != nil {
		return models.Usuario{}, err
	}
	defer rows.Close()
	var usuario models.Usuario
	if rows.Next() {
		if err := rows.Scan(&usuario.ID, &usuario.Senha, &usuario.Role, &usuario.Nome); err != nil {
			return models.Usuario{}, err
		}
		return usuario, nil
	}
	return models.Usuario{}, errors.New("Usuário não encontrado. Verifique as credenciais e tente novamente")
}

func (r RepoUsuarios) ProcuraPorID(ID int) (models.Usuario, error) {
	rows, err := r.db.Query("SELECT nome, nome_do_meio, ultimo_nome, nascimento, email, senha, role, numero_reservas, cadastrado FROM usuarios WHERE id = ?;", ID)
	if err != nil {
		return models.Usuario{}, err
	}
	defer rows.Close()

	var usuario models.Usuario
	if rows.Next() {
		if err := rows.Scan(&usuario.Nome,
			&usuario.NomeMeio,
			&usuario.NomeUltimo,
			&usuario.Nascimento,
			&usuario.Email,
			&usuario.Senha,
			&usuario.Role,
			&usuario.NumReservas,
			&usuario.CriadoEm,
		); err != nil {
			return models.Usuario{}, err
		}
	}
	return usuario, nil
}
