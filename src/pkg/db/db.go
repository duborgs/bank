package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB() (sql.DB, error) {

	//db, err := sql.Open("mysql", "root:root@tcp(localhost)/banco")
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3307)/q2bank")
	if err != nil {
		return *db, err
	}

	defer db.Close()

	return *db, nil
}
