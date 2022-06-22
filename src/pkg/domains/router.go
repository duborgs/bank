package main

import (
	"fmt"
	"pkg/domains/transaction/model"
	"pkg/domains/transaction/repository"
)

var transaction model.Transaction

func main() {
	transaction = model.Transaction{Value: "1000.0"}

	_, err := repository.InsertTransaction(transaction)
	if err != nil {
		fmt.Println(err)
	}

}
