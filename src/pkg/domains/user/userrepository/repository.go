package userRepository

import (
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

/*
MOCK
SE EXISTE OS DOIS USUARIOS
SALDO
TIPO DE USUARIO
*/
