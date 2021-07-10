package domain

import "github.com/djedjethai/bankingSqlx/errs"

type CustomerRepositoryStub struct {
	customers []Customer
}

func NewCustomerRepositoryStub() *CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "1"},
		{Id: "1002", Name: "Anna", City: "Bkk", ZipCode: "10110", DateOfBirth: "2002-2-2", Status: "0"},
	}

	return &CustomerRepositoryStub{customers}
}

func (s *CustomerRepositoryStub) FindAll(stt string) ([]Customer, *errs.AppError) {
	if stt == "" {
		return s.customers, nil
	} else {
		var listCust []Customer
		for i := range s.customers {
			if s.customers[i].Status == stt {
				listCust = append(listCust, s.customers[i])
			}
		}

		return listCust, nil
	}
}

func (s *CustomerRepositoryStub) ById(id string) (*Customer, *errs.AppError) {
	for i := range s.customers {
		if s.customers[i].Id == id {
			return &s.customers[i], nil
		}
	}

	return nil, errs.NewNotFoundError("Customer not found")
}
