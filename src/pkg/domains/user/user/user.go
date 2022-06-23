package user

import (
	"fmt"
)

type User struct {
	ID       int64
	Name     string
	Type     string
	CPF_CNPJ string
	Email    string
	Password string
	Wallet   float64
}

func FormatQueryUpdate(user User) string {
	query := fmt.Sprintf("update users set wallet = %v where id = %v ;", user.Wallet, user.ID)
	return query
}

func FormatQueryGet(ID int64) string {
	query := fmt.Sprintf("select * from users where id = %v", ID)
	return query
}
