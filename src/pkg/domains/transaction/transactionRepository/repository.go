package transactionRepository

import (
	"pkg/db"
	"pkg/domains/transaction/transaction"
)

func InsertTransaction(transactionInsert transaction.Transaction) error {

	db, err := db.OpenDB()
	if err != nil {
		return err
	}

	defer db.Close()

	query := transaction.FormatQueryInsert(transactionInsert)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
