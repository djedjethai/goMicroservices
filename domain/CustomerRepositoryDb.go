package domain

import (
	"database/sql"
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
