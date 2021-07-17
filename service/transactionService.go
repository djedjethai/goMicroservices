package service

import (
	"github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
	"time"
)

type TransactionService interface {
	HandleTransaction(dto.NewTransactionRequest) (dto.NewTransactionResponse, *errs.AppError)
}

type AccountRepository interface {
	GetBalance(string) (float64, *errs.AppError)
	UpdateAccountAmount(domain.Transaction) *errs.AppError
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

	// req account balance from account service
	balance, err := domain.GetBalance(t.AccountId)
	if err != nil {
		return nil, err
	}

	var newAmount float64
	if dt.TransactionType == "withdrawal" {
		if balance < t.Amount {
			return nil, errs.NewBadRequestError("Account found insufisant")
		}
		newAmount = balance - t.Amount
	} else if dt.TransactionType == "deposit" {
		newAmount = balance + t.Amount
	} else {
		return nil, errs.NewBadRequestError("wrong parameter")
	}

	if err := domain.UpdateAccountAmount(newAmount); err != nil {
		return nil, errs.NewInternalServerError("Unexpected database error")
	}

	dt := domain.Transaction{
		TransactionId:   "",
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transactionId, err := domain.UpdateTransactionTable(dt)
	if err != nil {
		return nil, errs.NewInternalServerError("Unexpected database error")
	}

	dtoResp := dto.NewTransactionResponse{
		Amount:        newAmout,
		TransactionId: transactionId,
	}

	return dtoResp, nil
}
