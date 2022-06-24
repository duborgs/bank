package main

import (
	"net/http"
	"pkg/domains/transaction/transport"
	userTransport "pkg/domains/user/transport"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/transaction", transport.TransactionHandle)
	http.HandleFunc("/getUserID", userTransport.GetUserIDHandler)
	http.HandleFunc("/getTransactions", transport.GetTransactionsHandler)
	http.ListenAndServe(":8080", nil)

}
