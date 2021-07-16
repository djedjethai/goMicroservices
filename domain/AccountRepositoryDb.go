package domain

import (
	"database/sql"
	"github.com/djedjethai/bankingSqlx/errs"
	"github.com/djedjethai/bankingSqlx/logger"
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

func (cl AccountRepositoryDb) GetBalance(an string) (float64, *errs.AppError) {
	balanceSql := "SELECT amount FROM accounts WHERE account_id = ?"

	var balance float64
	if err := cl.client.Get(&balance, balanceSql, an); err != nil {
		if err == sql.ErrNoRows {
			return balance, errs.NewBadRequestError("bad request from get balance")
		} else {
			logger.Error("err while scanning balance" + err.Error())
			return balance, errs.NewInternalServerError("Unexpected database err")
		}
	}

	return balance, nil
}

func (cl TransactionRepositoryDb) UpdateAccountAmount(aid string, a float64) *errs.AppError {
	updateSql := "UPDATE accounts SET amount = ? WHERE account_id = ?"

	_, err := cl.client.Exec(updateSql, a, aid)
	if err != nil {
		logger.Error("err while updating account" + err.Error())
		return errs.NewInternalServerError("Unexpected database error")
	}

	return nil
}
