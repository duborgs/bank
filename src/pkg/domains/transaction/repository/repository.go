package repository

import (
	"fmt"
	"pkg/db"
	"pkg/domains/transaction/transaction"
)

var (
	tr           transaction.Transaction
	transactions []transaction.Transaction
	query        string
)

func InsertTransaction(transactionInsert transaction.Transaction) error {

	db, err := db.OpenDB()
	if err != nil {
		return err
	}

	defer db.Close()

	query = transaction.FormatQueryInsert(transactionInsert)
	fmt.Print(query)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func GetTransactions(ID int64) ([]transaction.Transaction, error) {

	db, err := db.OpenDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	query = transaction.FormatQueryGet(ID)

	data, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for data.Next() {

		err = data.Scan(&tr.IDOrigin, &tr.IDDestin, &tr.Day_Time, &tr.Value)
		if err != nil {
			fmt.Printf("Error when performing the scan %v\n", err)
			return nil, err
		}
		transactions = append(transactions, tr)
	}

	return transactions, nil
}
