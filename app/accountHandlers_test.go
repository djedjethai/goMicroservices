package app

import (
	"github.com/djedjethai/bankingLib/errs"
	"github.com/djedjethai/bankingSqlx/dto"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_account_handler_create_new_account_return_an_error_if_body_in_req_is_not_jsonFormat(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	// set router handler
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.postAccount)
	// set req
	request, _ := http.NewRequest(http.MethodPost, "/customers/1001/account", strings.NewReader(""))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusBadRequest {
		t.Error("while testing create new account with no imput, err should not be null")
	}
}

func Test_account_handler_create_new_account_should_return_an_error_if_wrong_input(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// prepare http body content
	httpBody := `{
 		"account_type":"wronginput",
     		"amount": 6000
 	}`

	// set mocked service input
	mockSvcInput := dto.NewAccountRequest{
		CustomerId:  "1001",
		AccountType: "wronginput",
		Amount:      6000,
	}

	// set mock svc func
	mockServiceAccount.EXPECT().NewAccount(mockSvcInput).Return(nil, errs.NewValidationError("Account should be 'saving' or 'checking'"))

	// set router handler
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.postAccount)
	// set http req
	request, _ := http.NewRequest(http.MethodPost, "/customers/1001/account", strings.NewReader(httpBody))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusUnprocessableEntity {
		t.Error("While testing create new account with wrong input, err should not be null")
	}
}

func Test_account_handler_create_new_account_should_return_nil(t *testing.T) {

	tearDown := setup(t)
	defer tearDown()

	// prepare http body content
	bodyContent := `{
		"account_type":"saving",
    		"amount": 6000
	}`

	// set req passed to mocked service
	settedReq := dto.NewAccountRequest{
		CustomerId:  "1001",
		AccountType: "saving",
		Amount:      6000,
	}

	// set mocked service return expectation
	returnFromService := dto.NewAccountResponse{AccountId: "100"}

	// mock the service
	mockServiceAccount.EXPECT().NewAccount(settedReq).Return(&returnFromService, nil)

	// set the router handler
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.postAccount)
	// set the http req
	request, _ := http.NewRequest(http.MethodPost, "/customers/1001/account", strings.NewReader(bodyContent))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusCreated {
		t.Error("While testing route 'PostAccount', statusCode should be StatusCreated")
	}
}

// i am going to return an err from db
// but db is already tested by svc
// and svc return err will anyway be test if wrong input
// so i feel no need to test that .....??
// func Test_account_handler_create_new_account_should_return_an_error_if_wrong_wrong_input(t *testing.T) {
// 	tearDown := setup()
// 	defer tearDown()
//
// 	// prepare http body content
// 	httpBody := `{
// 		"account_type":"wronginput",
//     		"amount": 6000
// 	}`
//
// 	// set mocked service input
// 	mockSvcInput := dto.NewAccountRequest{
// 		CustomerId: "1001",
// 		AccountType: "saving",
// 		Amount: 6000,
// 	}
//
// 	// set mock svc func
// 	mockServiceAccount.EXPECT().NewAccount(mockSvcInput).Return(nil, errs.NewInternalServerError("service return an err"))
//
// 	//
// }
