package main

import (
	"pkg/domains/transaction/service"
	"pkg/domains/transaction/transaction"
)

func main() {
	var tr transaction.Transaction

	tr = transaction.Transaction{IDOrigin: 1, IDDestin: 2, Value: 0.1}

	service.MakeTransaction(tr)
}
