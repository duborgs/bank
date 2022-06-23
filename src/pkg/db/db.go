package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USER       = "ROOT"
	PASSWORD   = "123456"
	DB         = "Q2"
	PORT       = 3307
	PASSWORDDB = "123456"
)

func OpenDB() (sql.DB, error) {

	//db, err := sql.Open("mysql", "root:root@tcp(3306)/q2bank")
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3307)/q2bank")
	if err != nil {
		return *db, err
	}

	defer db.Close()

	return *db, nil
}
