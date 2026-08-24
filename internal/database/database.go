package database

import (
	"database/sql"
	"fmt"

	"gobook/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ConnString)
	if err != nil {
		fmt.Println(fmt.Errorf("erro ao abrir conexão SQL: %w", err))
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		fmt.Println(fmt.Errorf("erro ao abrir conexão SQL: %w", err))
		return nil, err
	}
	return db, nil
}
