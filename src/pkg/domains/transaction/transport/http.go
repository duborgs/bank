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
	request transaction.Transaction
	err     error
)

func TransactionHandle(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		fmt.Print("Metodo não é POST")
		return
	}

	request.IDOrigin, err = strconv.ParseInt(r.FormValue("payer"), 0, 64)
	if err != nil {
		log.Print("Erro ao realizar parse do IDOrigin")
		return
	}

	request.IDDestin, err = strconv.ParseInt(r.FormValue("payee"), 0, 64)
	if err != nil {
		log.Print("Erro ao realizar parse do IDOrigin")
		return
	}

	request.Value, err = strconv.ParseFloat(r.FormValue("value"), 64)
	if err != nil {
		log.Print("Erro ao realizar parse do IDOrigin")
		return
	}
	service.MakeTransaction(request)

	fmt.Printf("Bateu no handler")
}
