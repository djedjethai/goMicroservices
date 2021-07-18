package app

import (
	"encoding/json"
	// "fmt"
	"github.com/djedjethai/bankingSqlx/dto"
	// "github.com/djedjethai/bankingSqlx/errs"
	"github.com/djedjethai/bankingSqlx/service"
	// "github.com/gorilla/mux"
	"net/http"
)

type transactionHandlers struct {
	service service.TransactionService
}

func (s *transactionHandlers) postTransaction(w http.ResponseWriter, r *http.Request) {
	// get body
	var transactionReq dto.NewTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&transactionReq); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	dtoAccountResp, appErr := s.service.HandleTransaction(transactionReq)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, dtoAccountResp)
}
