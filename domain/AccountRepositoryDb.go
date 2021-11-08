package domain

import (
	"github.com/djedjethai/bankingLib/errs"
	"github.com/djedjethai/bankingLib/logger"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {

	return AccountRepositoryDb{dbClient}
}

func (cl AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {

	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := cl.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating an account: " + err.Error())
		return nil, errs.NewInternalServerError("Unexpected error from the database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while guetting last inserted id from inserted account: " + err.Error())
		return nil, errs.NewInternalServerError("Unexpected error from the database")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}
