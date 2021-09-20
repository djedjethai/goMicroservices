package app

import (
	"encoding/json"
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/service"
	"github.com/gorilla/mux"
	"net/http"
)

type accountHandlers struct {
	service service.AccountService
}

func (s accountHandlers) postAccount(w http.ResponseWriter, r *http.Request) {
	// extract body and params
	vars := mux.Vars(r)
	custId := vars["customer_id"]

	var na dto.NewAccountRequest
	if err := json.NewDecoder(r.Body).Decode(&na); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// add the customerId to na(new account) first
	na.CustomerId = custId

	accountResponse, appErr := s.service.NewAccount(na)
	if appErr != nil {
		writeResponse(w, appErr.Code, appErr.AsMessage())
	} else {
		writeResponse(w, http.StatusCreated, accountResponse)

	}
}
