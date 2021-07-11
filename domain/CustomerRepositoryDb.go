package domain

import (
	"database/sql"
	"fmt"
	"github.com/djedjethai/bankingSqlx/errs"
	"github.com/djedjethai/bankingSqlx/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {

	c := new(CustomerRepositoryDb)

	// this env var must be sanityCheck() (place it in app.go)
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	var err error
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	c.client, err = sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	c.client.SetConnMaxLifetime(time.Minute * 3)
	c.client.SetMaxOpenConns(10)
	c.client.SetMaxIdleConns(10)

	return c
}

func (cl *CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error

	customers := make([]Customer, 0)
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		// select from sqlx retreive the rows from db
		// it also marshall this retreive rows into our domain object
		err = cl.client.Select(&customers, findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = cl.client.Select(&customers, findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while quering cutomer table" + err.Error())
		return nil, errs.NewInternalServerError("Unexpected database error")
	}

	// no need anymore as we use Select which manage all
	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning cutomers" + err.Error())
	// 	return nil, errs.NewInternalServerError("Unexpected database error")

	// }

	return customers, nil
}

func (cl *CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {

	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	var c Customer
	if err := cl.client.Get(&c, customerSql, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning cutomers" + err.Error())
			return nil, errs.NewInternalServerError("Unexpected database error")
		}
	}

	return &c, nil
}
