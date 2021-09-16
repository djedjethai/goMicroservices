package app

import (
	"github.com/djedjethai/bankingSqlx/dto"
	// "github.com/djedjethai/bankingSqlx/errs"
	// "github.com/djedjethai/bankingSqlx/mocks/service"
	// "github.com/golang/mock/gomock"
	// "github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_account_handler_create_new_account_should_return_nil(t *testing.T) {
	// Arrange
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()
	// mockService = service.NewMockAccountService(ctrl)
	// ah = accountHandlers{mockService}

	tearDown := setup(t)
	defer tearDown()

	// prepare a fictif body content ??

	// router = mux.NewRouter()
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.postAccount)
	// make the http req
	request, _ := http.NewRequest(http.MethodPost, "/customers/{customer_id:[0-9]+}/account", nil)

	// set req to pass to service
	settedReq := dto.NewAccountRequest{
		CustomerId:  "1001",
		AccountType: "saving",
		Amount:      6000,
	}

	// set return from service
	returnFromService := dto.NewAccountResponse{AccountId: "100"}

	// mock the service
	mockService.EXPECT().NewAccount(settedReq).Return(returnFromService, nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder)

	// Assert
	if recorder.Code != http.StatusCreated {
		t.Error("While testing route 'PostAccount', statusCode should be StatusCreated")
	}
}
