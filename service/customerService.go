package service

import (
	"github.com/djedjethai/banking/domain"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer, error)
}

type Repository interface {
	FindAll() ([]domain.Customer, error)
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
