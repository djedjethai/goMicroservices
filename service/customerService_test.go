package service

import (
	realDomain "github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
	"testing"
)

// test GetAllCustomer
func Test_GetAllCustomer_should_return_from_db_allcustomers_if_argument_is_0(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	dummyCustomers := []realDomain.Customer{
		{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "0"},
	}
	expectedOutput := []dto.CustomerResponse{
		dto.CustomerResponse{
			Id:          dummyCustomers[0].Id,
			Name:        dummyCustomers[0].Name,
			City:        dummyCustomers[0].City,
			ZipCode:     dummyCustomers[0].ZipCode,
			DateOfBirth: dummyCustomers[0].DateOfBirth,
			Status:      dummyCustomers[0].Status,
		},
	}
	mockRepoCust.EXPECT().FindAll("0").Return(dummyCustomers, nil)

	// Act
	outputFromService, err := custService.GetAllCustomer("inactive")

	// Assert
	if err != nil {
		t.Error("While testing 'GetAllCustomer(1)', err should be null")
	}
	if expectedOutput[0].Id != outputFromService[0].Id {
		t.Error("while testing 'GetAllCustomer(1)', customer output have wrong id")
	}
}

func Test_GetAllCustomer_should_return_from_db_allcustomers_if_argument_is_1(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	// mock func FindAll() should return
	dummyCustomers := []realDomain.Customer{
		{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "1"},
	}

	// set the expected output from GetAllCustomers()
	var expectedOutput []dto.CustomerResponse
	expectedOutput = append(expectedOutput, dto.CustomerResponse{
		Id:          dummyCustomers[0].Id,
		Name:        dummyCustomers[0].Name,
		City:        dummyCustomers[0].City,
		ZipCode:     dummyCustomers[0].ZipCode,
		DateOfBirth: dummyCustomers[0].DateOfBirth,
		Status:      dummyCustomers[0].Status,
	},
	)

	// set the mock func
	mockRepoCust.EXPECT().FindAll("1").Return(dummyCustomers, nil)

	// Act
	custFromService, err := custService.GetAllCustomer("active")

	// Assert
	if err != nil {
		t.Error("While testing 'GetAllCustomer(1)', err should be null")
	}
	if custFromService[0].Id != expectedOutput[0].Id {
		t.Error("while testing 'GetAllCustomer(1)', customer output have wrong id")
	}
}

func Test_GetAllCustomer_should_return_from_db_allcustomers_if_argument_is_empty(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	dummyCustomers := []realDomain.Customer{
		{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "1"},
		{Id: "1002", Name: "Anna", City: "Bkk", ZipCode: "10110", DateOfBirth: "2002-2-2", Status: "0"},
	}

	// create the expected result from function: dummyCustomers to dto...
	var dtoResponse []dto.CustomerResponse
	for _, cust := range dummyCustomers {
		dtoResponse = append(dtoResponse, dto.CustomerResponse{
			Id:          cust.Id,
			Name:        cust.Name,
			City:        cust.City,
			ZipCode:     cust.ZipCode,
			DateOfBirth: cust.DateOfBirth,
			Status:      cust.Status,
		})
	}

	mockRepoCust.EXPECT().FindAll("").Return(dummyCustomers, nil)

	// Act
	custFromDb, err := custService.GetAllCustomer("")

	// Assert
	if err != nil {
		t.Error("While testing 'GetAllCustomer()', err should be null")
	}

	// THE COMPARAISON OF ID SHOULD BE DONE WITH THE CREATED EXPECTED RES(dto....)
	if custFromDb[0].Id != dtoResponse[0].Id {
		t.Error("while testing 'GetAllCustomer()', customer output have wrong id")
	}
	if custFromDb[1].Id != dtoResponse[1].Id {
		t.Error("while testing 'GetAllCustomer()', customer have wrong id")
	}
}

// test Getcustomer method
func Test_GetCustomer_should_return_a_database_error(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	mockRepoCust.EXPECT().ById("1001").Return(nil, errs.NewInternalServerError("unexpected db err"))

	// Act
	outputFromService, err := custService.GetCustomer("1001")

	// Assert
	if err == nil {
		t.Error("While testing 'GetCustomer' db err, err should not be null")
	}
	if outputFromService != nil {
		t.Error("While testing 'GetCustomer' db err, output shoud be null")
	}
}

func Test_GetCustomer_should_return_selected_customer(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	dummyCustomers := realDomain.Customer{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "1"}

	mockRepoCust.EXPECT().ById("1001").Return(&dummyCustomers, nil)

	// Act
	outputFromService, err := custService.GetCustomer("1001")

	// Assert
	if err != nil {
		t.Error("While testing 'GetCustomer' err should be null")
	}
	if dummyCustomers.Name != outputFromService.Name {
		t.Error("While testing 'GetCustomer' name shoud match")
	}
}
