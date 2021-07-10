package service

import (
	"github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
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

func (s *defaultRepositoryService) GetAllCustomer(status string) ([]dto.CustomerResponse, *errs.AppError) {
	// 'active' or 'inactive' is the value of the query 'status'
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	var dtoc []dto.CustomerResponse
	lc, err := s.repos.FindAll(status)
	if err != nil {
		return dtoc, err
	}

	for _, c := range lc {
		cust := c.ToDto()
		dtoc = append(dtoc, cust)
	}

	return dtoc, nil
}

func (s *defaultRepositoryService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repos.ById(id)
	if err != nil {
		return nil, err
	}

	cr := c.ToDto()

	return &cr, nil
}
