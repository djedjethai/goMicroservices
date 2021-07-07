package domain

import (
	"database/sql"
	"github.com/djedjethai/banking/errs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func NewCustomerRepositoryDb() *CustomerRepositoryDb {

	c := new(CustomerRepositoryDb)

	var err error
	c.client, err = sql.Open("mysql", "root:root@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	c.client.SetConnMaxLifetime(time.Minute * 3)
	c.client.SetMaxOpenConns(10)
	c.client.SetMaxIdleConns(10)

	return c
}

func (cl *CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := cl.client.Query(findAllSql)
	if err != nil {
		log.Println("Error while quering cutomer table" + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		if err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status); err != nil {
			log.Println("Error while scanning cutomers" + err.Error())
			return nil, err

		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (cl *CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {

	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := cl.client.QueryRow(customerSql, id)

	var c Customer
	if err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status); err != nil {
		// manage the err, this method is from sql module
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			log.Println("Error while scanning cutomers" + err.Error())
			return nil, errs.NewInternalServerError("Unexpected database error")
		}
	}

	return &c, nil
}
