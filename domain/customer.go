package domain

import (
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
)

type Customer struct {
	Id          string `db:"customer_id"` // for sqlx package to marshall with the db matchingName
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.statusAsText(),
	}
}
