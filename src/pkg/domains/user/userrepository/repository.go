package userRepository

import (
	"fmt"
	"pkg/db"
	"pkg/domains/user/user"

	_ "github.com/go-sql-driver/mysql"
)

func UpsertUser(userUp user.User) error {

	db, err := db.OpenDB()
	if err != nil {
		return err
	}

	query := user.FormatQueryUpdate(userUp)

	_, err = db.Exec(query)
	if err != nil {

		return err
	}

	db.Close()

	return nil
}

func GetUserID(ID int64) (user.User, error) {

	var userID user.User

	query := user.FormatQueryGet(ID)

	db, err := db.OpenDB()
	if err != nil {
		fmt.Printf("Error opening Data Base	%v\n", err)
		return userID, nil
	}

	defer db.Close()

	date, err := db.Query(query)
	if err != nil {
		fmt.Printf("Error executing query %v\n", err)
		return user.User{}, nil
	}

	date.Next()
	err = date.Scan(&userID.ID, &userID.Name, &userID.Type, &userID.CPF_CNPJ, &userID.Email, &userID.Password, &userID.Wallet)
	if err != nil {
		fmt.Printf("Error when performing the scan %v\n", err)
		return user.User{}, nil
	}

	return userID, nil
}

/*
MOCK
SE EXISTE OS DOIS USUARIOS
SALDO
TIPO DE USUARIO
*/
