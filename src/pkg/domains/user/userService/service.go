package userService

import (
	"encoding/json"
	"pkg/domains/user/userRepository"
	"strconv"
)

var (
	IDConv int64
	err    error
)

func GetUserID(ID string) (string, error) {

	if ID == "" {
		return "", err
	}

	IDConv, err = strconv.ParseInt(ID, 0, 64)
	if err != nil {
		return "Error", err
	}

	userResponse, err := userRepository.GetUserID(IDConv)
	if err != nil {
		return "Error", err
	}
	userResponse.Password = ""

	resp, err := json.Marshal(userResponse)

	return string(resp), nil
}
