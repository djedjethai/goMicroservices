package app

import (
	"encoding/json"
	"fmt"
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
	"github.com/djedjethai/bankingSqlx/service"
	"github.com/gorilla/mux"
	"net/http"
)

type transactionHandlers struct {
	service service.transactionServ
}

func (s *transactionHandlers) postTransaction(w http.ResponseWriter, r *http.Request) {
	// get param
	vars := mux.Vars(r)
	accountId := vars["account_id"]

	// get body
	var transactionReq dto.NewTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&transactionReq); err != nil {
		err.writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	transactionReq.AccountId = accountId

	dtoAccountResp, appErr := s.service.HandleTransaction(transactionReq)
	if err != nil {
		writeResponse(w, AppErr.Code, appErr.AsMessage())
		return
	}

	writeResponse(w, http.StatusOk, dtoAccountResp)
}
