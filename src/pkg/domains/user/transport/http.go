package transport

import (
	"fmt"
	"net/http"
	"pkg/domains/user/service"
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

	resp, err = service.GetUserID(ID)
	if err != nil {
		fmt.Print("Erro no getuser", err)
		return
	}
	fmt.Fprintf(w, "%v", resp)

}
