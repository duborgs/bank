package repository

import (
	"pkg/db"
	"pkg/domains/transaction/model"
)

func InsertTransaction(transaction model.Transaction) (string, error) {

	db, err := db.OpenDB()
	if err != nil {
		return "Erro na conexao com banco", err
	}

	defer db.Close()

	query := model.FormatQuery(transaction)

	_, err = db.Exec(query)
	if err != nil {
		return "Erro ao realizar query", err
	}

	return "Transação concluida", nil
}
