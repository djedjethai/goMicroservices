package service

import (
	"github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
	"time"
)

type TransactionService interface {
	HandleTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type TransactionRepository interface {
	RunTransaction(domain.Transaction, float64) (*dto.NewTransactionResponse, *errs.AppError)
	GetBalance(string) (float64, *errs.AppError)
}

type transactionService struct {
	service TransactionRepository
}

func NewTransactionService(transacDb TransactionRepository) TransactionService {
	return &transactionService{transacDb}
}

func (s *transactionService) HandleTransaction(t dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	// get account amount
	balance, errB := s.service.GetBalance(t.AccountId)
	if errB != nil {
		return nil, errB
	}

	if err := t.Validate(balance); err != nil {
		return nil, err
	}

	dt := domain.Transaction{
		TransactionId:   "",
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	dtoResp, err := s.service.RunTransaction(dt, balance)
	if err != nil {
		return nil, errs.NewInternalServerError("Unexpected database error")
	}

	return dtoResp, nil
}
