package main

import (
	"net/http"
	"pkg/domains/transaction/transport"
)

func main() {
	// var tr transaction.Transaction

	// tr = transaction.Transaction{IDOrigin: 1, IDDestin: 2, Value: 4000}

	// service.MakeTransaction(tr)

	http.HandleFunc("/transaction", transport.TransactionHandle)
	http.ListenAndServe(":8080", nil)
}
