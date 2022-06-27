package service

import (
	"encoding/json"
	"pkg/domains/user/model"
	"pkg/domains/user/repository"
	"strconv"
)

var (
	IDConv       int64
	err          error
	responseB    []byte
	responseS    string
	userResponse model.User
)

func GetUserID(ID string) (string, error) {

	if ID == "" {
		return "", err
	}

	IDConv, err = strconv.ParseInt(ID, 0, 64)
	if err != nil {
		return "Error in convert ID", err
	}

	userResponse, err = repository.GetUserID(IDConv)
	if err != nil {
		return "Error when search user", err
	}

	userResponse.Password = ""

	responseB, err = json.Marshal(userResponse)

	return string(responseB), nil
}

func CreateUser(userCreate model.User) (string, error) {

	switch {
	case userCreate.Type == "commom":
		break
	case userCreate.Type == "shopkeeper":
		break
	default:
		return "Invalid type", err
	}

	if userCreate.Wallet < 0 {
		return "Wallet cannot be negative", err
	}

	responseS, err = repository.CreateUser(userCreate)
	if err != nil {
		return responseS, err
	}
	return "User registered successfully", err
}
