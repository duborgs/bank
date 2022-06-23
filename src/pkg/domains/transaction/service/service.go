package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
		return "Erro na busca do mock", err
	}

	if mock != true {
		return "Serviço indisponivel", err
	}

	payer, err := user.GetUserID(MTransaction.IDOrigin)
	if err != nil {
		log.Fatal("Payer não existe")
	}

	payee, err := user.GetUserID(MTransaction.IDDestin)
	if err != nil {
		log.Fatal("Payee não existe")
	}

	token, cond := ValidatePayer(payer, MTransaction.Value)
	if token != true {
		log.Fatalf("Permissão negada: a%v", cond)
	}

	err = transactionRepository.InsertTransaction(MTransaction)
	if err != nil {
		fmt.Print("Erro ao inserir transação")
		log.Fatal(err)
	}

	payer.SALDO = payer.SALDO - MTransaction.Value

	payee.SALDO = payee.SALDO + MTransaction.Value

	err = userRepository.UpsertUser(payer)

	if err != nil {
		fmt.Print("Erro ao atualizar saldo")
		log.Fatal(err)
	}

	err = userRepository.UpsertUser(payee)
	if err != nil {
		fmt.Print("Erro ao atualizar saldo")
		log.Fatal(err)
	}

	return "Nada", nil
}

func ValidatePayer(payer user.User, Wallet float64) (bool, string) {

	if payer.TIPO != "comum" {
		return false, "Somente lojistas podem fazer transferencia"
	}

	if payer.SALDO < Wallet {
		return false, "Saldo insuficiente"
	}

	return true, ""
}

func GetMock() (bool, error) {
	var url URL

	fmt.Print(url)
	fmt.Print("\n")
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
