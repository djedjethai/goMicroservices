package service

import (
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
)

type TransactionService interface {
	HandleTransaction() (dto.NewTransactionResponse, *errs.AppError)
}

type TransactionRepository interface {
	RunTransaction()
}

type AccountRepository interface {
	// add this 2 method to account serviceDb
	GetBalance(string) (float64, errs.AppError)
	UpdateAccountAmount(float64) errs.AppError
}

type transactionService struct {
	transactionDb TransactionRepository
}

func NewTransactionService(transacDb domain.Transaction) *TransactionService {
	return &transactionService{transacDb}
}

func (s *transactionService) HandleTransaction(t dto.NewTransactionRequest) (*dto.NewAccountResponse, *errs.AppError) {

	// check if withdrawal or deposit
	if 

	// req account balance from account service

	// make sure the account amount(from accountdb) is ok


	// transform type from dto to Transaction
	// update transaction table in transaction service
	// return transactionId

	// calcule the new amount and save to accountdb

	// return updatedBalance + transactionId
}
