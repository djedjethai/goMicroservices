package domain

import (
	"github.com/djedjethai/bankingLib/errs"
	"github.com/djedjethai/bankingSqlx/dto"
)

type Transaction struct {
	TransactionId   string `json:"transaction_id"` // for sqlx to match to db column
	AccountId       string `json:"account_id"`
	Amount          float64
	TransactionType string `json:"transaction_type"`
	TransactionDate string `json:"transaction_date"`
}

//go:generate mockgen -destination=../mocks/domain/mockTransactionRepository.go -package=domain github.com/djedjethai/bankingSqlx/domain TransactionRepository
// run the mock: go generate ./...
type TransactionRepository interface {
	RunTransaction(Transaction, float64) (*dto.NewTransactionResponse, *errs.AppError)
	GetBalance(string) (float64, *errs.AppError)
}

func (t Transaction) CanNotWithdraw(amount float64) bool {
	if t.Amount < amount {
		return true
	}

	return false
}
