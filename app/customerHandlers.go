package app

import (
	"encoding/json"
	// "encoding/xml"
	"github.com/djedjethai/bankingLib/logger"
	"github.com/djedjethai/bankingSqlx/service"
	"github.com/gorilla/mux"
	"net/http"
)

// type Customer struct {
// 	Name    string `json:"full_name" xml:"name"`
// 	City    string `json:"city" xml:"city"`
// 	ZipCode string `json:"zip_code" xml:"zipcode"`
// }

type customerHandlers struct {
	service service.CustomerService
}

func (s *customerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// get query
	status := r.URL.Query().Get("status")

	customers, err := s.service.GetAllCustomer(status)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}

	// if r.Header.Get("content-type") == "application/xml" {
	// 	w.Header().Add("Content-type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)
	// } else {
	// 	w.Header().Add("Content-type", "application/json") // otherwise its text/plain
	// 	json.NewEncoder(w).Encode(customers)
	// }
}

func (s *customerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// parse the segment(the params..)
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := s.service.GetCustomer(id)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {

	// this should always be in this order
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		logger.Error("Error when writer encode data")
		panic(err)
	}
}

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "post req received")
// }
