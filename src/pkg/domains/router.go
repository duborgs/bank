package main

import (
	"net/http"
	"pkg/domains/transaction/transactionTransport"
	"pkg/domains/user/userTransport"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	http.HandleFunc("/transaction", transactionTransport.TransactionHandle)
	http.HandleFunc("/GetUserID", userTransport.GetUserIDHandler)
	http.ListenAndServe(":8080", nil)

}
