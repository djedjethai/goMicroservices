package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/djedjethai/banking/service"
	"net/http"
	"strconv"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zipcode"`
}

type customerHandlers struct {
	service service.CustomerService
}

func (s *customerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := s.service.GetAllCustomer()

	if r.Header.Get("content-type") == "application/xml" {
		w.Header().Add("Content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-type", "application/json") // otherwise its text/plain
		json.NewEncoder(w).Encode(customers)
	}
}

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "hello World")
// }

func (s *customerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	// parse the segment(the params..)
	vars := mux.Vars(r)
	id := vars["customer_id"]

	idInt, _ := strconv.Itoa()

	customer, err := s.ById(idInt)
	if err != nil {
		w.StatusCode(http.StatusNotFound)
		fmt.Fprintf(w, err.Error())
	} else {
		w.Header().Add("Content-type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "post req received")
// }
