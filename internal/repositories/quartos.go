package repositories

import (
	"database/sql"
	"gobook/internal/models"
)

type RepoQuartos struct {
	db *sql.DB
}

func NewQuartosRepo(db *sql.DB) *RepoQuartos {
	return &RepoQuartos{db}
}

func (r RepoQuartos) BuscarTodosQuartosDeUsuarioPorID(userID int) ([]models.Quarto, error) {
	query := `
		SELECT id, nome, descricao, valor_noite, reembolso, foto, disponivel, propriedade, dono
		FROM quartos 
		WHERE dono = ?
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quartos []models.Quarto

	for rows.Next() {
		var q models.Quarto
		err := rows.Scan(
			&q.ID,
			&q.Nome,
			&q.Descricao,
			&q.ValorNoite,
			&q.Reembolso,
			&q.Foto,
			&q.Disponivel,
			&q.Propriedade,
			&q.Dono,
		)
		if err != nil {
			return nil, err
		}
		quartos = append(quartos, q)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return quartos, nil
}

func (r RepoQuartos) BuscarTodosQuartosDePropriedadePorID(propriedadeID int) ([]models.Quarto, error) {
	query := `
		SELECT id, nome, descricao, valor_noite, reembolso, foto, disponivel, propriedade, dono
		FROM quartos
		WHERE propriedade = ?
	`
	rows, err := r.db.Query(query, propriedadeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var quartos []models.Quarto
	for rows.Next() {
		var q models.Quarto
		if err := rows.Scan(
			&q.ID,
			&q.Nome,
			&q.Descricao,
			&q.ValorNoite,
			&q.Reembolso,
			&q.Foto,
			&q.Disponivel,
			&q.Propriedade,
			&q.Dono,
		); err != nil {
			return nil, err
		}
		quartos = append(quartos, q)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return quartos, nil
}

func (r RepoQuartos) CriaQuarto(q models.Quarto, userID int) error {
	query := `
		INSERT INTO quartos (nome, descricao, valor_noite, reembolso, foto, disponivel, propriedade, dono)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query,
		q.Nome,
		q.Descricao,
		q.ValorNoite,
		q.Reembolso,
		q.Foto,
		q.Disponivel,
		q.Propriedade,
		userID,
	)

	return err
}

func (r RepoQuartos) BuscaQuartoPorNome(nome string, userID int) (models.Quarto, error) {
	query := `
		SELECT id, nome, descricao, valor_noite, reembolso, foto, disponivel, propriedade, dono
		FROM quartos 
		WHERE nome = ? AND dono = ?
	`
	var q models.Quarto

	err := r.db.QueryRow(query, nome, userID).Scan(
		&q.ID,
		&q.Nome,
		&q.Descricao,
		&q.ValorNoite,
		&q.Reembolso,
		&q.Foto,
		&q.Disponivel,
		&q.Propriedade,
		&q.Dono,
	)

	if err != nil {
		return models.Quarto{}, err
	}

	return q, nil
}

func (r RepoQuartos) DeletaQuartoPorNome(nome string, userID int) error {
	query := `DELETE FROM quartos WHERE nome = ? AND dono = ?`

	_, err := r.db.Exec(query, nome, userID)

	return err
}

func (r RepoQuartos) AtualizaQuarto(q models.Quarto) error {
	query := `
		UPDATE quartos 
		SET nome = ?, descricao = ?, valor_noite = ?, reembolso = ?, foto = ?, disponivel = ?
		WHERE id = ? AND dono = ?
	`

	_, err := r.db.Exec(query,
		q.Nome,
		q.Descricao,
		q.ValorNoite,
		q.Reembolso,
		q.Foto,
		q.Disponivel,
		q.ID,
		q.Dono,
	)

	return err
}
