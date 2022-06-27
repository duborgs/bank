package main

import (
	"net/http"
	"pkg/domains/transaction/transport"
	userTransport "pkg/domains/user/transport"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//Router transaction
	http.HandleFunc("/transaction", transport.TransactionHandler)
	http.HandleFunc("/getTransactions", transport.GetTransactionsHandler)

	//Router user
	http.HandleFunc("/getUserID", userTransport.GetUserIDHandler)
	http.HandleFunc("/createUser", userTransport.CreateUserHandler)

	http.ListenAndServe(":8080", nil)

}
