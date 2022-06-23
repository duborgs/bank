package transaction

import "fmt"

type Transaction struct {
	Value    float64
	IDOrigin int64
	IDDestin int64
	Day      string
}

func FormatQueryInsert(transaction Transaction) string {
	query := fmt.Sprintf("INSERT INTO TRANSACTIONS (Amount, ID_ORIGIN, ID_DEST) VALUES (%v, %v, %v)", transaction.Value, transaction.IDOrigin, transaction.IDDestin)
	return query
}
