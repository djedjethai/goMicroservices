package domain

import (
	"database/sql"
	"github.com/djedjethai/bankingSqlx/errs"
	"github.com/djedjethai/bankingSqlx/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	client *sql.DB
}

func NewTransactionRepositoryDb(client *sql.DB) TransactionRepository {
	return TransactionRepository{client}
}

func (tr TransactionRepository) RunTransaction() (transactionId, errs.AppError) {
	// add the Transaction
}
