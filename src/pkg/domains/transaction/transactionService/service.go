package transactionService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pkg/domains/transaction/transaction"
	"pkg/domains/transaction/transactionRepository"
	"pkg/domains/user/user"
	"pkg/domains/user/userRepository"
)

type URL struct {
	Authorization bool `json:"authorization"`
}

func GetTransactionID(ID user.User) {

}

func MakeTransaction(MTransaction transaction.Transaction) (string, error) {

	/*DECODE DO MOCK*/

	mock, err := GetMock()
	if err != nil {
		return "Error fetching mock", err
	}

	if mock != true {
		return "Unavailable service", err
	}

	payer, err := userRepository.GetUserID(MTransaction.IDOrigin)
	if err != nil {
		return "Payer does not exist", err
	}

	payee, err := userRepository.GetUserID(MTransaction.IDDestin)
	if err != nil {
		return "Payee does not exist", err
	}

	token, cond := ValidatePayer(payer, MTransaction.Value)
	if token != true {

		validadeMSG := fmt.Sprintf("Permission denied: %v", cond)
		return validadeMSG, err
	}

	err = transactionRepository.InsertTransaction(MTransaction)
	if err != nil {
		return "Error entering transaction in Data Base", err
	}

	payer.Wallet -= MTransaction.Value

	payee.Wallet += MTransaction.Value

	err = userRepository.UpsertUser(payer)

	if err != nil {
		return "Error updating payer wallet", err
	}

	err = userRepository.UpsertUser(payee)
	if err != nil {
		return "Error updating payee wallet", err
	}

	return "Transaction completed successfully", nil
}

func ValidatePayer(payer user.User, Wallet float64) (bool, string) {

	if payer.Type != "commom" {
		return false, "Shopkeeper cannot transfer"
	}

	if payer.Wallet < Wallet {
		return false, "Insufficient amount in wallet"
	}

	return true, ""
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
