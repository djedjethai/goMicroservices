package domain

import (
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	GetBalance(string) (float64, *errs.AppError)
	UpdateAccountAmount(Transaction) *errs.AppError
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {

	return dto.NewAccountResponse{AccountId: a.AccountId}
}
