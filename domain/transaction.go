package domain

import "github.com/djedjethai/bankingSqlx/errs"

type Transaction struct {
	TransactionId   string `json:"transaction_id"` // for sqlx to match to db column
	AccountId       string `json:"account_id"`
	Amount          float64
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
}

type TransactionRepository interface {
	UpdateTransactionTable(Transaction) (string, *errs.AppError)
	GetBalance(string) (float64, *errs.AppError)
	UpdateAccountAmount(float64, string) *errs.AppError
}
