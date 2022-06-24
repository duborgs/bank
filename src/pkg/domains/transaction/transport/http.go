package transport

import (
	"fmt"
	"log"
	"net/http"
	"pkg/domains/transaction/service"
	"pkg/domains/transaction/transaction"
	"strconv"
)

var (
	request           transaction.Transaction
	err               error
	msg, response, ID string
	msgErro           int
)

func TransactionHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Fprint(w, "Method is not POST")
		return
	}

	request.IDOrigin, err = strconv.ParseInt(r.FormValue("payer"), 0, 64)
	if err != nil {
		fmt.Fprint(w, "Error converting payer ID")
		return
	}

	request.IDDestin, err = strconv.ParseInt(r.FormValue("payee"), 0, 64)
	if err != nil {
		fmt.Fprint(w, "Error converting payee ID")
		return
	}

	if request.IDOrigin == request.IDDestin {
		fmt.Fprint(w, "The payer cannot transfer to himself")
		return
	}
	request.Value, err = strconv.ParseFloat(r.FormValue("value"), 64)
	if err != nil {
		log.Print("Erro ao realizar parse do IDOrigin")
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
		fmt.Print(err)
		return
	}
	if response == "" {
		fmt.Print(err)
		return
	}

	fmt.Fprint(w, response)
}
