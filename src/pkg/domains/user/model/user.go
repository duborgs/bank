package model

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

var (
	query string
)

func FormatQueryUpdate(user User) string {
	query = fmt.Sprintf("update users set wallet = %v where id = %v ;", user.Wallet, user.ID)
	return query
}

func FormatQueryGet(ID int64) string {
	query = fmt.Sprintf("select * from users where id = %v", ID)
	return query
}

func FormatQueryCreate(user User) string {
	user.Type = fmt.Sprint("'" + user.Type + "'")

	query = fmt.Sprintf("insert into users (name_user, type_user, CPF_CNPJ, email, password_user, wallet) values (%v, %v, %v, %v, %v, %v);", user.Name, user.Type, user.CPF_CNPJ, user.Email, user.Password, user.Wallet)
	return query
}
