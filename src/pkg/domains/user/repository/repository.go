package repository

import (
	"pkg/db"
	"pkg/domains/user/model"

	_ "github.com/go-sql-driver/mysql"
)

var (
	err       error
	query, er string
)

func UpsertWallet(userUp model.User) error {

	db, err := db.OpenDB()
	if err != nil {
		return err
	}
	defer db.Close()
	query := model.FormatQueryUpdate(userUp)

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func GetUserID(ID int64) (model.User, error) {

	var userID model.User

	query = model.FormatQueryGet(ID)

	db, err := db.OpenDB()
	if err != nil {
		return userID, err
	}

	defer db.Close()

	date, err := db.Query(query)
	if err != nil {
		return model.User{}, err
	}

	date.Next()
	err = date.Scan(&userID.ID, &userID.Name, &userID.Type, &userID.CPF_CNPJ, &userID.Email, &userID.Password, &userID.Wallet)
	if err != nil {
		return model.User{}, err
	}

	return userID, nil
}

func CreateUser(user model.User) (string, error) {

	db, err := db.OpenDB()
	if err != nil {
		return "Error opening Data Base", err
	}

	defer db.Close()

	query = model.FormatQueryCreate(user)

	_, err = db.Query(query)
	if err != nil {
		return "Error executing query", err
	}
	return "Successful query", nil

}
