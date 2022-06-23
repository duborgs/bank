package transaction

import (
	"fmt"
	"time"
)

type Transaction struct {
	Value    float64
	IDOrigin int64
	IDDestin int64
	Day_Time string
}

func FormatQueryInsert(transaction Transaction) string {
	dateFormat := fmt.Sprintf("'%v'", time.Now().Format("2006-01-02 15:04:05"))
	transaction.Day_Time = dateFormat

	query := fmt.Sprintf("INSERT INTO TRANSACTIONS (Amount, ID_ORIGIN, ID_DEST, DATE_TIME) VALUES (%v, %v, %v, %v)", transaction.Value, transaction.IDOrigin, transaction.IDDestin, transaction.Day_Time)
	return query
}
