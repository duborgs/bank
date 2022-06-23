package transactionTransport

import (
	"fmt"
	"log"
	"net/http"
	"pkg/domains/transaction/transaction"
	"pkg/domains/transaction/transactionService"
	"strconv"
)

var (
	request transaction.Transaction
	err     error
	msg     string
	msgErro int
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

	msg, err = transactionService.MakeTransaction(request)
	//fmt.Fprintf(w, "%v", msg)

}
