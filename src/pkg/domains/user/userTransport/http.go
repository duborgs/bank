package userTransport

import (
	"fmt"
	"net/http"
	"pkg/domains/user/userService"
)

var (
	ID, resp, msg string
	err           error
	msgErro       int
)

func GetUserIDHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		fmt.Fprint(w, "Method is not POST")
		return
	}

	ID = r.FormValue("ID")

	resp, err = userService.GetUserID(ID)
	if err != nil {
		fmt.Print("Erro no getuser", err)
		return
	}
	fmt.Fprintf(w, "%v", resp)

}
