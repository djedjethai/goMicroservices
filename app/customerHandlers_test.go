package app

import (
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"testing"
)

func Test_should_return_customers_with_status_code_200(t *testing.T) {
	// Arrange
	// the controller manage the state for the mock
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // !!! important
	// create the mock service
	mockService := service.NewMockCustomerService(ctrl)
	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "1"},
		{Id: "1002", Name: "Anna", City: "Bkk", ZipCode: "10110", DateOfBirth: "2002-2-2", Status: "0"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
	// define customer handler
	// inject the mock service
	ch := customerHandlers{mockService}

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.getAllCustomers)

	// make the http req
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	// a finir .......
}
