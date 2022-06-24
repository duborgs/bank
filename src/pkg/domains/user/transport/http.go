package transport

import (
	"fmt"
	"net/http"
	"pkg/domains/user/model"
	"pkg/domains/user/service"
	"strconv"
)

var (
	ID, response, msg string
	err               error
	msgErro           string
	request           model.User
)

func GetUserIDHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		fmt.Fprint(w, "Method is not GET")
		return
	}

	ID = r.FormValue("ID")

	response, err = service.GetUserID(ID)
	if err != nil {
		fmt.Print("Erro in getuser", err)
		return
	}
	fmt.Fprintf(w, "%v", response)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	msgErro = "Value not null"

	if r.Method != http.MethodPost {
		fmt.Fprint(w, "Method is not POST")
		return
	}

	if request.Name = r.FormValue("name_user"); request.Name == "" {
		return
	}

	if request.Type = r.FormValue("type_user"); request.Type == "" {
		return
	}

	if request.CPF_CNPJ = r.FormValue("cpf_cnpj"); request.CPF_CNPJ == "" {
		return
	}

	if request.Email = r.FormValue("email"); request.Email == "" {
		return
	}

	if request.Password = r.FormValue("password"); request.Password == "" {
		return
	}

	request.Wallet, err = strconv.ParseFloat(r.FormValue("wallet"), 64)
	if err != nil {
		fmt.Fprint(w, "Erro in conversion")
		return
	}

	response, err = service.CreateUser(request)
	if err != nil {
		fmt.Fprint(w, response)
		return
	}

	fmt.Fprint(w, response)

}
