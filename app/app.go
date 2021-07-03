package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	// mux := http.NewServeMux() // standart http multiplexer
	router := mux.NewRouter()

	// method masher
	router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	// req masher customer_id must be int
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
