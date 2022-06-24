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
	response     []byte
	userResponse model.User
)

func GetUserID(ID string) (string, error) {

	if ID == "" {
		return "", err
	}

	IDConv, err = strconv.ParseInt(ID, 0, 64)
	if err != nil {
		return "Error", err
	}

	userResponse, err = repository.GetUserID(IDConv)
	if err != nil {
		return "Error", err
	}

	userResponse.Password = ""

	response, err = json.Marshal(userResponse)

	return string(response), nil
}
