package model

import "fmt"

type User struct {
	NOME  string
	TIPO  string
	CPF   string
	CNPJ  string
	EMAIL string
	SENHA string
	SALDO string
}

type CreateUserRequest struct {
	ID    string
	SALDO string
}

func FormatQuery(user CreateUserRequest) string {
	query := fmt.Sprintf("UPDATE USUARIOS SET SALDO = %s WHERE ID = %s ;", user.SALDO, user.ID)
	return query
}
