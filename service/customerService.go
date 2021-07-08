package service

import (
	"github.com/djedjethai/banking/domain"
	"github.com/djedjethai/banking/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type Repository interface {
	FindAll() ([]domain.Customer, *errs.AppError)
	ById(string) (*domain.Customer, *errs.AppError)
}

type defaultRepositoryService struct {
	repos Repository
}

func NewService(repos Repository) CustomerService {
	return &defaultRepositoryService{repos}
}

func (s *defaultRepositoryService) GetAllCustomer() ([]domain.Customer, *errs.AppError) {
	return s.repos.FindAll()
}

func (s *defaultRepositoryService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repos.ById(id)
}
