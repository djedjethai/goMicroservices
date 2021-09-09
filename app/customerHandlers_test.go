package app

import (
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
	"github.com/djedjethai/bankingSqlx/mocks/service"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *mux.Router
var mockService *service.MockCustomerService
var ch customerHandlers

func setup(t *testing.T) func() {
	// Arrange
	// the controller manage the state for the mock
	ctrl := gomock.NewController(t)
	// create the mock service
	mockService = service.NewMockCustomerService(ctrl)

	// define customer handler
	// inject the mock service
	ch = customerHandlers{mockService}

	router = mux.NewRouter()
	// router.HandleFunc("/customers", ch.getAllCustomers)

	// close the ctrl as cb func
	// otherwise it will close the connection before the test finished
	return func() {
		router = nil
		defer ctrl.Finish() // !!! important
	}
}

func Test_should_return_error_message_with_status_code_500_for_selected_customer(t *testing.T) {
	// Arrange
	tearDown := setup(t)
	defer tearDown()

	// set an unexisting id
	mockService.EXPECT().GetCustomer("2006").Return(nil, errs.NewInternalServerError("some err in db"))

	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer)
	request, _ := http.NewRequest(http.MethodGet, "/customers/2006", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("While testing 'getCustomer', statusCode should be 500")
	}
}

func Test_should_return_selected_customers_with_status_code_200(t *testing.T) {
	// Arrange
	tearDown := setup(t)
	defer tearDown()

	dummyCustomer := &dto.CustomerResponse{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "1"}
	mockService.EXPECT().GetCustomer("1001").Return(dummyCustomer, nil)

	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer)
	request, _ := http.NewRequest(http.MethodGet, "/customers/1001", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("While testing 'getCustomer' the statusCode should be 200")
	}
}

func Test_should_return_customers_with_status_code_200(t *testing.T) {

	tearDown := setup(t)
	defer tearDown()

	dummyCustomers := []dto.CustomerResponse{
		{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "1"},
		{Id: "1002", Name: "Anna", City: "Bkk", ZipCode: "10110", DateOfBirth: "2002-2-2", Status: "0"},
	}
	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)

	// make the http req
	router.HandleFunc("/customers", ch.getAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	// recorder is an implementation of ResponseWriter(for the test)
	recorder := httptest.NewRecorder()
	// this code will send the req to the end-point
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Testing '/customers' route, failed when testing the statusOK")
	}
}

func Test_should_return_error_message_with_status_code_500(t *testing.T) {
	// Arrange
	tearDown := setup(t)
	defer tearDown()
	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewInternalServerError("some db errs"))
	router.HandleFunc("/customers", ch.getAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Testing '/customer' route, failed while testing StatusInternalServerError")
	}
}

// func Test_should_return_customers_with_status_code_200(t *testing.T) {
// 	// Arrange
// 	// the controller manage the state for the mock
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish() // !!! important
// 	// create the mock service
// 	mockService := service.NewMockCustomerService(ctrl)
// 	dummyCustomers := []dto.CustomerResponse{
// 		{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "1"},
// 		{Id: "1002", Name: "Anna", City: "Bkk", ZipCode: "10110", DateOfBirth: "2002-2-2", Status: "0"},
// 	}
// 	mockService.EXPECT().GetAllCustomer("").Return(dummyCustomers, nil)
//
// 	// define customer handler
// 	// inject the mock service
// 	ch := customerHandlers{mockService}
// 	router := mux.NewRouter()
// 	router.HandleFunc("/customers", ch.getAllCustomers)
// 	// make the http req
// 	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
//
// 	// Act
// 	// recorder is an implementation of ResponseWriter(for the test)
// 	recorder := httptest.NewRecorder()
// 	// this code will send the req to the end-point
// 	router.ServeHTTP(recorder, request)
//
// 	// Assert
// 	if recorder.Code != http.StatusOK {
// 		t.Error("Testing '/customers' route, failed when testing the statusOK")
// 	}
// }
//
// func Test_should_return_error_message_with_status_code_500(t *testing.T) {
// 	// Arrange
// 	// the controller manage the state for the mock
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish() // !!! important
// 	// create the mock service
// 	mockService := service.NewMockCustomerService(ctrl)
//
// 	mockService.EXPECT().GetAllCustomer("").Return(nil, errs.NewInternalServerError("some db errs"))
//
// 	// define customer handler
// 	// inject the mock service
// 	ch := customerHandlers{mockService}
// 	router := mux.NewRouter()
// 	router.HandleFunc("/customers", ch.getAllCustomers)
// 	// make the http req
// 	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
//
// 	// Act
// 	recorder := httptest.NewRecorder()
// 	router.ServeHTTP(recorder, request)
//
// 	// Assert
// 	if recorder.Code != http.StatusInternalServerError {
// 		t.Error("Testing '/customer' route, failed while testing StatusInternalServerError")
// 	}
// }
