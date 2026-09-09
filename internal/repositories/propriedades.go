package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"gobook/internal/models"
)

// Como essa parte da query é repetida algumas vezes, declarei aqui para não repetir esse texto enorme toda hora
var selectPropriedades = "SELECT propriedades.id, propriedades.nome, propriedades.endereco, propriedades.numero, propriedades.descricao, propriedades.estado, propriedades.cidade, propriedades.pet_friendly, propriedades.categoria, propriedades.dono_id, usuarios.nome, usuarios.nome_do_meio, usuarios.ultimo_nome FROM propriedades left join usuarios on usuarios.id = propriedades.dono_id"

type RepoPropriedades struct {
	db *sql.DB
}

func NewPropriedadesRepo(db *sql.DB) *RepoPropriedades {
	return &RepoPropriedades{db}
}

func (r RepoPropriedades) BuscarTodasPropriedades() ([]models.Propriedade, error) {
	rows, err := r.db.Query(
		selectPropriedades + ";")
	if err != nil {
		return []models.Propriedade{}, err
	}
	defer rows.Close()

	var propriedades []models.Propriedade
	for rows.Next() {
		var propriedade models.Propriedade
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Endereco, &propriedade.Numero, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.PetFriendly, &propriedade.Categoria, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
		propriedades = append(propriedades, propriedade)
	}
	return propriedades, nil
}

func (r RepoPropriedades) BuscarDezPrimeirasPropriedadesAleatorio() ([]models.Propriedade, error) {
	rows, err := r.db.Query(
		selectPropriedades + " order by rand() limit 8")
	if err != nil {
		return []models.Propriedade{}, err
	}
	defer rows.Close()

	var propriedades []models.Propriedade
	for rows.Next() {
		var propriedade models.Propriedade
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Endereco, &propriedade.Numero, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.PetFriendly, &propriedade.Categoria, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
		propriedades = append(propriedades, propriedade)
	}
	return propriedades, nil
}

func (r RepoPropriedades) BuscarTodasPorDono(userID int) ([]models.Propriedade, error) {
	rows, err := r.db.Query(
		selectPropriedades+" WHERE propriedades.dono_id = ?;", userID)
	if err != nil {
		return []models.Propriedade{}, err
	}
	defer rows.Close()

	var propriedades []models.Propriedade
	for rows.Next() {
		var propriedade models.Propriedade
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Endereco, &propriedade.Numero, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.PetFriendly, &propriedade.Categoria, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
		propriedades = append(propriedades, propriedade)
	}
	return propriedades, nil
}

func (r RepoPropriedades) BuscaPropriedadePorNome(nome string) (models.Propriedade, error) {
	rows, err := r.db.Query(
		selectPropriedades+" WHERE propriedades.nome = ?;", nome)
	if err != nil {
		return models.Propriedade{}, err
	}

	var propriedade models.Propriedade
	if rows.Next() {
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Endereco, &propriedade.Numero, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.Categoria, &propriedade.PetFriendly, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
		return propriedade, nil
	}
	return models.Propriedade{}, errors.New("propriedade inexistente")
}

func (r RepoPropriedades) BuscaPropriedadePorID(ID string) (models.Propriedade, error) {
	rows, err := r.db.Query(
		selectPropriedades+" WHERE propriedades.id = ?;", ID)
	if err != nil {
		return models.Propriedade{}, err
	}
	defer rows.Close()

	var p models.Propriedade
	if rows.Next() {
		rows.Scan(&p.ID, &p.Nome, &p.Descricao, &p.Estado, &p.Cidade, &p.PetFriendly, &p.Categoria, &p.Dono.ID, &p.Dono.Nome, &p.Dono.NomeMeio, &p.Dono.NomeUltimo)
	}
	return p, nil
}

func (r RepoPropriedades) CriaPropriedade(p *models.Propriedade, donoID int) (string, error) {
	stmt, err := r.db.Prepare(
		"INSERT into propriedades (nome, descricao, endereco, numero, estado, cidade, pet_friendly, categoria, dono_id) values (?, ?, ?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Nome, p.Descricao, p.Endereco, p.Numero, p.Estado, p.Cidade, p.PetFriendly, p.Categoria, donoID)
	if err != nil {
		return "", err
	}

	return p.Nome, nil
}

func (r RepoPropriedades) VerificaDono(propriedadeID, userID int) (string, error) {
	rows, err := r.db.Query(
		"SELECT nome, dono_id FROM propriedades WHERE id = ?", propriedadeID,
	)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var (
		donoPropID int
		nome       string
	)
	if rows.Next() {
		if err := rows.Scan(&nome, &donoPropID); err != nil {
			return "", err
		}
		if donoPropID == userID {
			return nome, nil
		}
		fmt.Println(userID, donoPropID)

		return "", errors.New("a propriedade não existe ou você não tem permissão para acessá-la")
	}
	return "", errors.New("a query não trouxe nenhum resultado")
}

func (r RepoPropriedades) RetornaNome(propID int) (string, error) {
	var nome string
	rows, err := r.db.Query(
		"SELECT nome FROM propriedades WHERE id = ?", propID,
	)
	if err != nil {
		return "", nil
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&nome); err != nil {
			return "", err
		}
	}
	return nome, nil
}

func (r RepoPropriedades) DeletaPropriedade(ID int) error {
	stmt, err := r.db.Prepare(
		"DELETE FROM propriedades WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()
	sqlResponse, err := stmt.Exec(ID)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if rowsAff, _ := sqlResponse.RowsAffected(); rowsAff == 0 {
		return errors.New("a propriedade não existe ou você não tem permissão para acessá-la")
	}

	return nil
}

func (r RepoPropriedades) EditaPropriedade(p models.Propriedade, propID int) error {
	stmt, err := r.db.Prepare(
		"UPDATE propriedades SET nome = ?, descricao = ?, estado = ?, cidade = ?, pet_friendly = ?, categoria = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	sqlResponse, err := stmt.Exec(p.Nome, p.Descricao, p.Estado, p.Cidade, p.PetFriendly, p.Categoria, propID)
	if err != nil {
		return err
	}

	if rowsAff, _ := sqlResponse.RowsAffected(); rowsAff == 0 {
		return errors.New("a propriedade não existe ou você não tem permissão para acessá-la")
	}
	return nil
}
