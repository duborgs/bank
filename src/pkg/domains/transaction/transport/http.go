package transport

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"pkg/domains/transaction/service"
	"pkg/domains/transaction/transaction"
)

var (
	request           transaction.Transaction
	err               error
	msg, response, ID string
	msgErro           int
)

func TransactionHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprint(w, "Method is not POST")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprint(w, "Error converting json", err)
	}
	err = json.Unmarshal(body, &request)
	if err != nil {
		fmt.Fprint(w, "Error when ", err)
	}

	if request.IDDestin < 0 {
		fmt.Fprint(w, "Error in id payer")
		return
	}

	if request.IDOrigin < 0 {
		fmt.Fprintf(w, "Error in id payee")
		return

	}

	msg, err = service.MakeTransaction(request)

	fmt.Fprintf(w, "%v", msg)

}

func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		fmt.Fprint(w, "Method is not GET")
		return
	}

	ID = r.FormValue("ID")

	response, err = service.GetTransactions(ID)
	if err != nil {
		fmt.Fprintf(w, "%v", err)
		return
	}

	if response == "" {
		fmt.Fprintf(w, "%v", err)
		return
	}

	fmt.Fprintf(w, "%v", response)
}
