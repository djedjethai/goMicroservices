package app

import (
	"fmt"
	"github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWD") == "" ||
		os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("environment variable not define")
	}
}

func Start() {

	//to make sure the env var are here
	sanityCheck()

	// wiring alternatively, the Stub or the db
	// repos := domain.NewCustomerRepositoryStub()
	repos := domain.NewCustomerRepositoryDb()

	serv := service.NewService(repos)

	ch := customerHandlers{serv}

	// mux := http.NewServeMux() // standart http multiplexer
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	// the regex make sure only int can be passed as param, but are string when mux extract
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)
	// method masher
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// req masher customer_id must be int
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}
