package domain

import (
	"database/sql"
	"github.com/djedjethai/bankingSqlx/errs"
	"github.com/djedjethai/bankingSqlx/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
)

type TransactionRepositoryDb struct {
	client *sqlx.DB
}

func NewTransactionRepositoryDb(dbClient *sqlx.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{dbClient}
}

func (db TransactionRepositoryDb) UpdateTransactionTable(t Transaction) (string, *errs.AppError) {
	transactionSql := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) values (?, ?, ?, ?)"

	var transactionId string
	result, err := db.client.Exec(transactionSql, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		logger.Error("Error while updating transaction table" + err.Error())
		return transactionId, errs.NewInternalServerError("Unexpected server error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while guetting last transactionId" + err.Error())
		return transactionId, errs.NewInternalServerError("Unexpected server error")
	}

	transactionId = strconv.FormatInt(id, 10)
	return transactionId, nil
}

func (db TransactionRepositoryDb) GetBalance(an string) (float64, *errs.AppError) {
	balanceSql := "SELECT amount FROM accounts WHERE account_id = ?"

	var balance float64
	if err := db.client.Get(&balance, balanceSql, an); err != nil {
		if err == sql.ErrNoRows {
			return balance, errs.NewBadRequestError("bad request from get balance")
		} else {
			logger.Error("err while scanning balance" + err.Error())
			return balance, errs.NewInternalServerError("Unexpected database err")
		}
	}

	return balance, nil
}

func (db TransactionRepositoryDb) UpdateAccountAmount(amt float64, aid string) *errs.AppError {
	updateSql := "UPDATE accounts SET amount = ? WHERE account_id = ?"

	_, err := db.client.Exec(updateSql, amt, aid)
	if err != nil {
		logger.Error("err while updating account" + err.Error())
		return errs.NewInternalServerError("Unexpected database error")
	}

	return nil
}
