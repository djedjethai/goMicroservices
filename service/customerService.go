package service

import (
	"github.com/djedjethai/banking/domain"
	"github.com/djedjethai/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer, *errs.AppError)
	GetCustomer(string) (*domain.Customer, *errs.AppError)
}

type Repository interface {
	FindAll(string) ([]domain.Customer, *errs.AppError)
	ById(string) (*domain.Customer, *errs.AppError)
}

type defaultRepositoryService struct {
	repos Repository
}

func NewService(repos Repository) CustomerService {
	return &defaultRepositoryService{repos}
}

func (s *defaultRepositoryService) GetAllCustomer(status string) ([]domain.Customer, *errs.AppError) {
	// 'active' or 'inactive' is the value of the query 'status'
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	return s.repos.FindAll(status)
}

func (s *defaultRepositoryService) GetCustomer(id string) (*domain.Customer, *errs.AppError) {
	return s.repos.ById(id)
}
