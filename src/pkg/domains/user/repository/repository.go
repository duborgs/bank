package repository

import (
	"fmt"
	"pkg/db"
	"pkg/domains/user/model"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err   error
	query string
)

func UpsertWallet(userUp model.User) error {

	db, err := db.OpenDB()
	if err != nil {
		return err
	}

	query := model.FormatQueryUpdate(userUp)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	db.Close()

	return nil
}

func GetUserID(ID int64) (model.User, error) {

	var userID model.User

	query = model.FormatQueryGet(ID)

	db, err := db.OpenDB()
	if err != nil {
		fmt.Printf("Error opening Data Base	%v\n", err)
		return userID, nil
	}

	defer db.Close()

	date, err := db.Query(query)
	if err != nil {
		fmt.Printf("Error executing query %v\n", err)
		return model.User{}, nil
	}

	date.Next()
	err = date.Scan(&userID.ID, &userID.Name, &userID.Type, &userID.CPF_CNPJ, &userID.Email, &userID.Password, &userID.Wallet)
	if err != nil {
		fmt.Printf("Error when performing the scan %v\n", err)
		return model.User{}, nil
	}

	return userID, nil
}

func CreateUser(user model.User) (string, error) {

	db, err := db.OpenDB()
	if err != nil {
		return "Error opening Data Base", err
	}

	defer db.Close()

	fmt.Print(user)

	query = model.FormatQueryCreate(user)
	fmt.Print(query)
	_, err = db.Query(query)
	if err != nil {
		return "Error executing query", err
	}
	return "Successful query", nil

}
