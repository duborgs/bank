package transaction

import (
	"fmt"
	"time"
)

type Transaction struct {
	Value    float64 `json:"value"`
	IDOrigin int64   `json:"payer"`
	IDDestin int64   `json:"payee"`
	Day_Time string  `json:"daytime,omitempty"`
}

func FormatQueryInsert(transaction Transaction) string {
	dateFormat := fmt.Sprintf("'%v'", time.Now().Format("2006-01-02 15:04:05"))
	transaction.Day_Time = dateFormat

	query := fmt.Sprintf("insert into transactions (amount, id_origin, id_dest, date_time) values (%v, %v, %v, %v)", transaction.Value, transaction.IDOrigin, transaction.IDDestin, transaction.Day_Time)
	return query
}

func FormatQueryGet(ID int64) string {

	query := fmt.Sprintf("select id_origin, id_dest, date_time, amount from transactions where id_origin = %v or id_dest = %v", ID, ID)
	return query
}
