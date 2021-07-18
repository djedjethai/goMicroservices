package dto

import (
	"github.com/djedjethai/bankingSqlx/errs"
)

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (nt *NewTransactionRequest) Validate(balance float64) *errs.AppError {
	if nt.Amount > balance {
		return errs.NewBadRequestError("Account found insufisant")
	}

	return nil
}
