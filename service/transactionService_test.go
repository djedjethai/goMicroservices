package service

import (
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
	"testing"
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

// func Test_transactionService_Handle_transaction_get_balance_return_balance (t *testing.T){}
// func Test_transactionService_Handle_transaction_run_transaction_return_err (t *testing.T){}
// func Test_transactionService_Handle_transaction_run_transaction_return_newTransactionResponse (t *testing.T){}
