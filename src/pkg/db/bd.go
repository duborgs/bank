package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
const (
	USER       = "ROOT"
	PASSWORD   = "123456"
	DB         = "Q2"
	PASSWORDDB = "123456"
)
*/

func OpenDB() (sql.DB, error) {

	db, err := sql.Open("mysql", "root:root@tcp(3306)/q2bank")
	if err != nil {
		return *db, err
	}
	defer db.Close()

	fmt.Print("Conexão Realizada")

	return *db, nil
}
