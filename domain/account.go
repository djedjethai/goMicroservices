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

//go:generate mockgen -destination=../mocks/domain/mockAccountRepository.go -package=domain github.com/djedjethai/bankingSqlx/domain AccountRepository
// run the mock: go generate ./...
type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {

	return dto.NewAccountResponse{AccountId: a.AccountId}
}
