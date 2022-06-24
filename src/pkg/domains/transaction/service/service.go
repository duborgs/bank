package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pkg/domains/transaction/repository"
	"pkg/domains/transaction/transaction"
	"pkg/domains/user/model"
	userRepository "pkg/domains/user/repository"
	"strconv"
)

type URL struct {
	Authorization bool `json:"authorization"`
}

func GetTransactionID(ID model.User) {

}

var (
	payer, payee        model.User
	err                 error
	mock                bool
	IDConv              int64
	transactionResponse []transaction.Transaction
	response            []byte
)

func MakeTransaction(MTransaction transaction.Transaction) (string, error) {

	/*DECODE DO MOCK*/

	mock, err = GetMock()
	if err != nil {
		return "Error fetching mock", err
	}

	if mock != true {
		return "Unavailable service", err
	}

	payer, err = userRepository.GetUserID(MTransaction.IDOrigin)
	if err != nil {
		return "Payer does not exist", err
	}

	payee, err = userRepository.GetUserID(MTransaction.IDDestin)
	if err != nil {
		return "Payee does not exist", err
	}

	token, cond := ValidatePayer(payer, MTransaction.Value)
	if token != true {

		validadeMSG := fmt.Sprintf("Permission denied: %v", cond)
		return validadeMSG, err
	}
	err = repository.InsertTransaction(MTransaction)
	if err != nil {
		return "Error entering transaction in Data Base", err
	}

	payer.Wallet -= MTransaction.Value

	payee.Wallet += MTransaction.Value

	err = userRepository.UpsertWallet(payer)
	if err != nil {
		return "Error updating payer wallet", err
	}

	err = userRepository.UpsertWallet(payee)
	if err != nil {
		return "Error updating payee wallet", err
	}

	return "Transaction completed successfully", nil
}

func ValidatePayer(payer model.User, Wallet float64) (bool, string) {

	if payer.Type != "commom" {
		return false, "Shopkeeper cannot transfer"
	}

	if payer.Wallet < Wallet {
		return false, "Insufficient amount in wallet"
	}

	return true, ""
}

func GetTransactions(ID string) (string, error) {

	IDConv, err = strconv.ParseInt(ID, 0, 64)
	if err != nil {
		return "Error in conversion:", err
	}

	transactionResponse, err = repository.GetTransactions(IDConv)
	if err != nil {
		return "Error fetching data from database:", err
	}

	response, err = json.Marshal(transactionResponse)
	if err != nil {
		return "Error when running marshal:", err
	}
	return string(response), nil
}

func GetMock() (bool, error) {

	var url URL

	response, err := http.Get("https://run.mocky.io/v3/d02168c6-d88d-4ff2-aac6-9e9eb3425e31")
	if err != nil {
		return false, err
	}
	resp, _ := ioutil.ReadAll(response.Body)

	err = json.Unmarshal(resp, &url)
	if err != nil {
		return url.Authorization, err
	}

	return url.Authorization, nil
}
