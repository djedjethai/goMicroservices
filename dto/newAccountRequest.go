package dto

import (
	"github.com/djedjethai/bankingLib/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (ar NewAccountRequest) Validate() *errs.AppError {
	if ar.Amount < 5000 {
		return errs.NewValidationError("To open an account 5000$ is require")
	}

	if strings.ToLower(ar.AccountType) != "saving" && strings.ToLower(ar.AccountType) != "checking" {
		return errs.NewValidationError("Account should be 'saving' or 'checking'")
	}

	return nil
}
