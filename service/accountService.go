package service

import (
	"github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
	"time"
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

	a := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	newAccount, err := s.repos.Save(a)
	if err != nil {
		return nil, err
	}

	accountResp := newAccount.ToNewAccountResponseDto()

	return &accountResp, nil
}
