package model

import "fmt"

type Transaction struct {
	Value    string
	IDOrigin string
	IDDestin string
	Day      string
}

func FormatQuery(transaction Transaction) string {
	query := fmt.Sprintf("INSERT INTO TRANSACTION (Amount, ID_ORIGIN, ID_DEST) VALUES (%s, %s, %s)", transaction.Value, transaction.IDOrigin, transaction.IDDestin)
	return query
}
