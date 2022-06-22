package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pkg/domains/transaction/model"
)

type URL struct {
	Authorization bool `json:"authorization"`
}

func MakeTransaction(transaction model.Transaction) (string, error) {

	/*DECODE DO MOCK*/

	mock, err := SearchMock()
	if err != nil {
		return "Erro na busca do mock", err
	}

	if mock != true {
		return "Serviço indisponivel", err
	}

	fmt.Print(mock)

	return "Decode concluido", nil
}

func SearchMock() (bool, error) {
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
