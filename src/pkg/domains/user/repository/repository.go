package repository

import (
	"database/sql"
	"fmt"
	"pkg/domains/user/model"

	_ "github.com/go-sql-driver/mysql"
)

func OpenBD() (sql.DB, error) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		return *db, err
	}
	return *db, nil
}

func UpsertUser(user model.CreateUserRequest) (string, error) {
	db, err := OpenBD()
	if err != nil {
		return "Erro na conexao com banco", err
	}

	query := model.FormatQuery(user)

	up, err := db.Exec("%s", query)
	if err != nil {
		return "Erro ao realizar UPSERT", nil
	}

	fmt.Print(up)

	db.Close()
	return "Transação concluida", nil
}

func GETUSER() {

}

/*
MOCK
SE EXISTE OS DOIS USUARIOS
SALDO
TIPO DE USUARIO
*/
