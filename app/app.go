package app

import (
	"github.com/djedjethai/banking/domain"
	"github.com/djedjethai/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Start() {

	// wiring
	repos := domain.NewCustomerRepositoryStub()

	serv := service.NewService(repos)

	ch := customerHandlers{serv}

	// mux := http.NewServeMux() // standart http multiplexer
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	// method masher
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// req masher customer_id must be int
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
