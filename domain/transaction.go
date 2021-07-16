package domain

import "github.com/djedjethai/bankingSqlx/errs"

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

type TransactionRepository interface {
	UpdateTransactionTable(Transaction) (string, errs.AppError)
}
