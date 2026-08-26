package repositories

import (
	"database/sql"
	"gobook/internal/models"
)

type RepoPropriedades struct {
	db *sql.DB
}

func NewPropriedadesRepo(db *sql.DB) *RepoPropriedades {
	return &RepoPropriedades{db}
}

func (r RepoPropriedades) BuscarTodasPropriedades() ([]models.Propriedade, error) {
	rows, err := r.db.Query(
		"SELECT propriedades.id, propriedades.nome, propriedades.descricao, propriedades.estado, propriedades.cidade, propriedades.pet_friendly, propriedades.categoria, propriedades.dono_id, usuarios.nome, usuarios.nome_do_meio, usuarios.ultimo_nome " +
			"FROM propriedades " +
			"left join usuarios " +
			"on usuarios.id = propriedades.dono_id;")
	if err != nil {
		return []models.Propriedade{}, err
	}
	defer rows.Close()

	var propriedades []models.Propriedade
	for rows.Next() {
		var propriedade models.Propriedade
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.Categoria, &propriedade.PetFriendly, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
		propriedades = append(propriedades, propriedade)
	}
	return propriedades, nil
}

func (r RepoPropriedades) BuscaPropriedadePorNome(nome string) (models.Propriedade, error) {
	rows, err := r.db.Query(
		"SELECT propriedades.id, propriedades.nome, propriedades.descricao, propriedades.estado, propriedades.cidade, propriedades.pet_friendly,categoria, propriedades.dono_id, usuarios.nome, usuarios.nome_do_meio, usuarios.ultimo_nome "+
			"FROM propriedades "+
			"left join usuarios "+
			"on usuarios.id = propriedades.dono_id WHERE propriedades.nome = ?;", nome)
	if err != nil {
		return models.Propriedade{}, err
	}

	var propriedade models.Propriedade
	if rows.Next() {
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.Categoria, &propriedade.PetFriendly, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
	}
	return propriedade, nil
}

func (r RepoPropriedades) BuscaPropriedadePorID(ID string) (models.Propriedade, error) {
	rows, err := r.db.Query(
		"SELECT propriedades.id, propriedades.nome, propriedades.descricao, propriedades.estado, propriedades.cidade, propriedades.pet_friendly,categoria, propriedades.dono_id, usuarios.nome, usuarios.nome_do_meio, usuarios.ultimo_nome "+
			"FROM propriedades "+
			"left join usuarios "+
			"on usuarios.id = propriedades.dono_id WHERE propriedades.id = ?;", ID)
	if err != nil {
		return models.Propriedade{}, err
	}

	var propriedade models.Propriedade
	if rows.Next() {
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.Categoria, &propriedade.PetFriendly, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
	}
	return propriedade, nil
}

func (r RepoPropriedades) CriaPropriedade(propriedade *models.Propriedade, donoID int) (string, error) {
	stmt, err := r.db.Prepare(
		"INSERT into propriedades (nome, descricao, estado, cidade, pet_friendly, categoria, dono_id) values (?, ?, ?, ?, ?, ?, ?)",
	)
	if err != nil {
		return "", err
	}

	_, err = stmt.Exec(propriedade.Nome, propriedade.Descricao, propriedade.Estado, propriedade.Cidade, propriedade.PetFriendly, propriedade.Categoria, donoID)
	if err != nil {
		return "", err
	}

	return propriedade.Nome, nil
}
