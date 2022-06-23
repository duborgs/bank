package service

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

func MakeTransaction(MTransaction transaction.Transaction) (string, error) {

	/*DECODE DO MOCK*/

	mock, err := GetMock()
	if err != nil {
		return "Erro ao buscar mock", err
	}

	if mock != true {
		return "Serviço indisponivel", err
	}

	payer, err := user.GetUserID(MTransaction.IDOrigin)
	if err != nil {
		return "Payer não existe", err
	}

	payee, err := user.GetUserID(MTransaction.IDDestin)
	if err != nil {
		return "Payee não existe", err
	}

	token, cond := ValidatePayer(payer, MTransaction.Value)
	if token != true {

		validadeMSG := fmt.Sprintf("Permissão negada: %v", cond)
		return validadeMSG, err
	}

	err = transactionRepository.InsertTransaction(MTransaction)
	if err != nil {
		return "Erro ao inserir transação no banco", err
	}

	payer.SALDO = payer.SALDO - MTransaction.Value

	payee.SALDO = payee.SALDO + MTransaction.Value

	err = userRepository.UpsertUser(payer)

	if err != nil {
		return "Erro ao atualizar saldo do payer", err
	}

	err = userRepository.UpsertUser(payee)
	if err != nil {
		return "Erro ao atualizar saldo do payee", err
	}

	return "Transaction completed successfully", nil
}

func ValidatePayer(payer user.User, Wallet float64) (bool, string) {

	if payer.TIPO != "comum" {
		return false, "Lojistas não podem fazer transferencia"
	}

	if payer.SALDO < Wallet {
		return false, "Saldo insuficiente"
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
