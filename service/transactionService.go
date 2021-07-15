package service

import (
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
)

type TransactionService interface {
	HandleTransaction() (dto.NewTransactionResponse, errs.AppError)
}

type TransactionRepository interface {
	RunTransaction()
}

type AccountRepository interface {
	// add this 2 method to account serviceDb
	GetAmount(string) float64
	UpdateAmount(float64) errs.AppError
}

type transactionService struct {
	transactionDb TransactionRepository
}

func NewTransactionService(transacDb domain.Transaction) *TransactionService {
	return &transactionService{transacDb}
}

func (s *transactionService) HandleTransaction() (dto.NewAccountResponse, errs.AppError) {

	// req account balance from account service

	// make sure the account amount(from accountdb) is enought

	// run transaction in db in transaction service
	// return transactionId

	// calcule the new amount and save to accountdb

	// return
}
