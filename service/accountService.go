package service

import (
	"github.com/djedjethai/bankingLib/errs"
	"github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/dto"
	// "time"
)

//go:generate mockgen -destination=../mocks/service/mockAccountService.go -package=service github.com/djedjethai/bankingSqlx/service AccountService
type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type AccountRepositoryDb interface {
	Save(domain.Account) (*domain.Account, *errs.AppError)
}

type defaultAccountService struct {
	// repos domain.AccountRepository
	repos AccountRepositoryDb
}

func NewAccountService(as AccountRepositoryDb) AccountService {
	return &defaultAccountService{as}
}

func (s *defaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	a := domain.CreateNewAccount(req.CustomerId, req.AccountType, req.Amount)

	if newAccount, err := s.repos.Save(a); err != nil {
		return nil, err
	} else {
		return newAccount.ToNewAccountResponseDto(), nil
	}
}
