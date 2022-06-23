package user

import (
	"fmt"
	"pkg/db"
)

type User struct {
	ID       int64
	NOME     string
	TIPO     string
	CPF_CNPJ string
	EMAIL    string
	SENHA    string
	SALDO    float64
}

func FormatQueryUpdate(user User) string {
	query := fmt.Sprintf("UPDATE users SET WALLET = %v WHERE ID = %v ;", user.SALDO, user.ID)
	return query
}

func FormatQueryGet(ID int64) string {
	query := fmt.Sprintf("SELECT * FROM users WHERE ID = %v", ID)
	return query
}

func GetUserID(ID int64) (User, error) {

	var user User

	query := FormatQueryGet(ID)

	db, err := db.OpenDB()
	if err != nil {
		fmt.Printf("Erro ao abrir banco %v\n", err)
		return User{}, nil
	}

	defer db.Close()

	date, err := db.Query(query)
	if err != nil {
		fmt.Printf("Erro ao executar query %v\n", err)
		return User{}, nil
	}

	date.Next()
	err = date.Scan(&user.ID, &user.NOME, &user.TIPO, &user.CPF_CNPJ, &user.EMAIL, &user.SENHA, &user.SALDO)
	if err != nil {
		fmt.Printf("Erro ao realizar scan %v\n", err)
		return User{}, nil
	}

	return user, nil
}
