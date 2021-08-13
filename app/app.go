package app

import (
	"fmt"
	"github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/service"
	// "github.com/djedjethai/bankingSqlx/logger"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"os"
	"time"
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

	dbClient := getDbClient()
	// wiring alternatively, the Stub or the db
	// repos := domain.NewCustomerRepositoryStub()
	customerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	transactionRepositoryDb := domain.NewTransactionRepositoryDb(dbClient)

	customerServ := service.NewService(customerRepositoryDb)
	accountServ := service.NewAccountService(accountRepositoryDb)
	transactionServ := service.NewTransactionService(transactionRepositoryDb)

	ch := customerHandlers{customerServ}
	ah := accountHandlers{accountServ}
	th := transactionHandlers{transactionServ}

	// mux := http.NewServeMux() // standart http multiplexer
	router := mux.NewRouter()
	router.
		HandleFunc("/customers", ch.getAllCustomers).
		Methods(http.MethodGet).
		Name("GetAllCustomers") // name to match the auth verification
	// the regex make sure only int can be passed as param, but are string when mux extract
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.
		HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.postAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.
		HandleFunc("/customers/transaction", th.postTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")
	// method masher
	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// req masher customer_id must be int
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	// middleware
	am := authMiddleware{domain.NewAuthRepository()}
	router.Use(am.authorizationHandler())

	// starting server
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))

}

func getDbClient() *sqlx.DB {
	// this env var must be sanityCheck() (place it in app.go)
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
