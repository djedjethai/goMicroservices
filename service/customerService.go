package service

import (
	"github.com/djedjethai/banking/domain"
	"github.com/djedjethai/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type Repository interface {
	FindAll() ([]domain.Customer, error)
	ById(string) (*domain.Customer, *errs.AppError)
}

type defaultRepositoryService struct {
	repos Repository
}

func NewService(repos Repository) CustomerService {
	return &defaultRepositoryService{repos}
}

func (s *defaultRepositoryService) GetAllCustomer() ([]domain.Customer, error) {
	return s.repos.FindAll()
}

func (s *defaultRepositoryService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repos.ById(id)
}
