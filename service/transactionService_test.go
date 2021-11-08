package service

import (
	"github.com/djedjethai/bankingLib/errs"
	"github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/dto"
	"testing"
	"time"
)

func Test_transactionService_Handle_transaction_get_balance_return_err(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	// set the mock transaction req
	transacReq := dto.NewTransactionRequest{
		AccountId:       "100",
		Amount:          3000,
		TransactionType: "withdrawal",
	}

	mockRepoTransac.EXPECT().GetBalance(transacReq.AccountId).Return((float64)(0), errs.NewBadRequestError("Unknow bank account"))

	// Act
	_, err := transacService.HandleTransaction(transacReq)

	// Assert
	if err == nil {
		t.Error("while testing transactionService HandleTransaction get balance should return an err")
	}
}

func Test_transactionService_Handle_transaction_run_transaction_return_err(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	// set dto.transactionRequest for GetBalance
	transacReq := dto.NewTransactionRequest{
		AccountId:       "1001",
		Amount:          1000,
		TransactionType: "withdrawal",
	}

	// set domain.Transaction for RunTransaction
	transacInput := domain.Transaction{
		AccountId:       transacReq.AccountId,
		Amount:          transacReq.Amount,
		TransactionType: transacReq.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	mockRepoTransac.EXPECT().GetBalance(transacReq.AccountId).Return((float64)(5500.50), nil)

	mockRepoTransac.EXPECT().RunTransaction(transacInput, 5500.50).Return(nil, errs.NewInternalServerError("database unexpected error"))

	// Act
	_, err := transacService.HandleTransaction(transacReq)

	// Assert
	if err == nil {
		t.Error("while testing transactionService runTransaction should return an error ")
	}
}

func Test_transactionService_Handle_transaction_run_transaction_return_NewTransactionResponse(t *testing.T) {
	tearDown := setup(t)
	defer tearDown()

	// Arrange
	// set dto.transactionRequest for GetBalance
	transacReq := dto.NewTransactionRequest{
		AccountId:       "1001",
		Amount:          1000,
		TransactionType: "withdrawal",
	}

	// set domain.Transaction for RunTransaction
	transacInput := domain.Transaction{
		AccountId:       transacReq.AccountId,
		Amount:          transacReq.Amount,
		TransactionType: transacReq.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	// set dto.NewTransactionResponse
	transacResponse := dto.NewTransactionResponse{
		Amount:        3000,
		TransactionId: "234",
	}

	mockRepoTransac.EXPECT().GetBalance(transacReq.AccountId).Return((float64)(5500.50), nil)

	mockRepoTransac.EXPECT().RunTransaction(transacInput, 5500.50).Return(&transacResponse, nil)

	// Act
	responseFromService, _ := transacService.HandleTransaction(transacReq)

	// Assert
	if transacResponse.TransactionId != responseFromService.TransactionId {
		t.Error("while testing transactionService runTransaction should return an error ")
	}

}
