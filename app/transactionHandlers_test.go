package app

import (
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_transactionHandler_postTransaction_return_err_if_input_is_not_json(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	// set routeHandler and http req
	router.HandleFunc("/customers/transaction", th.postTransaction)
	request, _ := http.NewRequest(http.MethodPost, "/customers/transaction", strings.NewReader(""))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusBadRequest {
		t.Error("While testing TransactionHandler with text body error should be StatusBadRequest")
	}

}

func Test_transactionHandler_postTransaction_return_err_if_service_HandleTransaction_return_err(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	// set dto.NewTransactionRequest
	transacReq := dto.NewTransactionRequest{
		AccountId:       "1001",
		Amount:          500,
		TransactionType: "deposit",
	}

	// set json req
	jsonReq := `{"account_id":"1001", "amount":500, "transaction_type":"deposit"}`

	// set routeHandler and http req
	router.HandleFunc("/customers/transaction", th.postTransaction)
	request, _ := http.NewRequest(http.MethodPost, "/customers/transaction", strings.NewReader(jsonReq))

	mockServiceTransaction.EXPECT().HandleTransaction(transacReq).Return(nil, errs.NewInternalServerError("Unexpected database error"))

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("While testing transactionHandler, statusCode should return StatusInternalServerError")
	}
}

func Test_transactionHandler_postTransaction_return_statusCodeOK(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	//set dto.NewTransactionRequest
	transacReq := dto.NewTransactionRequest{
		AccountId:       "1001",
		Amount:          500,
		TransactionType: "withdrawal",
	}

	// set dto.NewTransactionResponse
	transacResp := dto.NewTransactionResponse{
		Amount:        5500,
		TransactionId: "123",
	}

	jsonReq := `{"account_id":"1001", "amount":500, "transaction_type":"withdrawal"}`

	// set router and request
	router.HandleFunc("/customers/transaction", th.postTransaction)
	request, _ := http.NewRequest(http.MethodPost, "/customers/transaction", strings.NewReader(jsonReq))

	mockServiceTransaction.EXPECT().HandleTransaction(transacReq).Return(&transacResp, nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("While testing transactionHandler, statusCode should return StatusOK")
	}
}
