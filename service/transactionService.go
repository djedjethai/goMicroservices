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
	UpdateTransactionTable(domain.Transaction) (string, *errs.AppError)
	GetBalance(string) (float64, *errs.AppError)
	UpdateAccountAmount(float64, string) *errs.AppError
}

type transactionService struct {
	service TransactionRepository
}

func NewTransactionService(transacDb TransactionRepository) *transactionService {
	return &transactionService{transacDb}
}

func (s *transactionService) HandleTransaction(t dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {

	// req account balance from account service
	balance, err := s.service.GetBalance(t.AccountId)
	if err != nil {
		return nil, err
	}

	var newAmount float64
	if t.TransactionType == "withdrawal" {
		if err := t.Validate(balance); err != nil {
			return nil, err
		}
		newAmount = balance - t.Amount
	} else if t.TransactionType == "deposit" {
		newAmount = balance + t.Amount
	} else {
		return nil, errs.NewBadRequestError("wrong parameter")
	}

	if err := s.service.UpdateAccountAmount(newAmount, t.AccountId); err != nil {
		return nil, errs.NewInternalServerError("Unexpected database error")
	}

	dt := domain.Transaction{
		TransactionId:   "",
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	transactionId, err := s.service.UpdateTransactionTable(dt)
	if err != nil {
		return nil, errs.NewInternalServerError("Unexpected database error")
	}

	dtoResp := dto.NewTransactionResponse{
		Amount:        newAmount,
		TransactionId: transactionId,
	}

	return &dtoResp, nil
}
