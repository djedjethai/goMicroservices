package domain

import (
	"database/sql"
	"github.com/djedjethai/bankingSqlx/errs"
	"github.com/djedjethai/bankingSqlx/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDb struct {
	client *sql.DB
}

func NewTransactionRepositoryDb(client *sql.DB) TransactionRepositoryDb {
	return TransactionRepositoryDb{client}
}

func (db TransactionRepositoryDb) UpdateTransactionTable(t Transaction) (string, errs.AppError) {
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
