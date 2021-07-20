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
	if nt.Amount < 0 {
		return errs.NewBadRequestError("Amount can not be negatif")
	}

	if nt.TransactionType == "withdrawal" && nt.Amount > balance {
		return errs.NewBadRequestError("Not enought found")
	}

	if nt.TransactionType != "deposit" && nt.TransactionType != "withdrawal" {
		return errs.NewBadRequestError("Unknow transaction type")
	}

	return nil
}
