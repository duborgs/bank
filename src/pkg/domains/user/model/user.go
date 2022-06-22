package model

import "fmt"

type User struct {
	ID    int64
	NOME  string
	TIPO  string
	CPF   string
	CNPJ  string
	EMAIL string
	SENHA string
	SALDO string
}

func FormatQuery(user User) string {
	query := fmt.Sprintf("UPDATE USUARIOS SET SALDO = %s WHERE ID = %v ;", user.SALDO, user.ID)
	return query
}

func GetUser() {

}

/*
BUSCA O PRIMEIRO USUARIO E SEGUNDO USUARIO
	VALIDA SE EXISTE
	O TIPO
	VALIDA O SALDO
*/
func GetType() {

}
