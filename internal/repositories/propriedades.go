package repositories

import (
	"gobook/internal/models"
)

func (r RepoUsuarios) BuscarTodasPropriedades() ([]models.Propriedade, error) {
	rows, err := r.db.Query(
		"SELECT propriedades.id, propriedades.nome, propriedades.descricao, propriedades.estado, propriedades.cidade, propriedades.categoria, propriedades.dono_id, usuarios.nome, usuarios.nome_do_meio, usuarios.ultimo_nome " +
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
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.Categoria, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
		propriedades = append(propriedades, propriedade)
	}
	return propriedades, nil
}

func (r RepoUsuarios) BuscaPropriedadePorNome(nome string) (models.Propriedade, error) {
	rows, err := r.db.Query(
		"SELECT propriedades.id, propriedades.nome, propriedades.descricao, propriedades.estado, propriedades.cidade, propriedades.categoria, propriedades.dono_id, usuarios.nome, usuarios.nome_do_meio, usuarios.ultimo_nome "+
			"FROM propriedades "+
			"left join usuarios "+
			"on usuarios.id = propriedades.dono_id WHERE propriedades.nome = ?;", nome)
	if err != nil {
		return models.Propriedade{}, err
	}

	var propriedade models.Propriedade
	if rows.Next() {
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.Categoria, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
	}
	return propriedade, nil
}

func (r RepoUsuarios) BuscaPropriedadePorID(ID string) (models.Propriedade, error) {
	rows, err := r.db.Query(
		"SELECT propriedades.id, propriedades.nome, propriedades.descricao, propriedades.estado, propriedades.cidade, propriedades.categoria, propriedades.dono_id, usuarios.nome, usuarios.nome_do_meio, usuarios.ultimo_nome "+
			"FROM propriedades "+
			"left join usuarios "+
			"on usuarios.id = propriedades.dono_id WHERE propriedades.id = ?;", ID)
	if err != nil {
		return models.Propriedade{}, err
	}

	var propriedade models.Propriedade
	if rows.Next() {
		rows.Scan(&propriedade.ID, &propriedade.Nome, &propriedade.Descricao, &propriedade.Estado, &propriedade.Cidade, &propriedade.Categoria, &propriedade.Dono.ID, &propriedade.Dono.Nome, &propriedade.Dono.NomeMeio, &propriedade.Dono.NomeUltimo)
	}
	return propriedade, nil
}
