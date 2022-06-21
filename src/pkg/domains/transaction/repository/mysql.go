package repository

import (
	"database/sql"
	"pkg/domains/user/model"

	_ "github.com/go-sql-driver/mysql"
)

func OpenBD() (sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return *db, err
	}
	return *db, nil
}

func UpsertTransaction(model.User) (string, error) {
	db, err := OpenBD()
	if err != nil {
		return "Erro na conexao com banco", err
	}

	query, err := FormatQuery(model.User)

	up, err := db.Exec("%s", query)
	if err != nil {
		return "Erro ao realizar UPDATE", error
	}

	db.Close()
	return "Transação concluida", nil
}
