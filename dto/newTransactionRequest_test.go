package dto

import (
	"net/http"
	"testing"
)

// test transaction type
func Test_should_return_error_when_transaction_type_is_not_deposit_or_withdrawl(t *testing.T) {
	// Arrange
	request := NewTransactionRequest{
		TransactionType: "invalid transaction type",
	}

	// Act
	AppError := request.Validate(123)

	// Assert message
	if AppError.Message != "Unknow transaction type" {
		t.Error("Invalid message while testing transaction type")
	}
	// Assert code
	if AppError.Code != http.StatusBadRequest {
		t.Error("Invalid err statusCode while testing transaction type")
	}
}
func Test_should_not_return_error_as_transaction_type_is_deposit(t *testing.T) {
	// Arrange
	request := NewTransactionRequest{
		TransactionType: "deposit",
	}

	// Act
	AppError := request.Validate(123)

	// Assert message
	if AppError != nil {
		t.Error("test deposit transaction, should not return error")
	}
}

// test withdraw amount and TransactionType as "withdrawal"
func Test_account_withdraw_is_negatif(t *testing.T) {
	request := NewTransactionRequest{
		TransactionType: "withdrawal",
		Amount:          100,
	}

	AppError := request.Validate(80)

	if AppError.Message != "Not enought found" {
		t.Error("Invalid err message while testing withdraw account")
	}

	if AppError.Code != http.StatusBadRequest {
		t.Error("Invalid err code while testing withdraw account")
	}
}
func Test_account_withdraw_is_positif(t *testing.T) {
	request := NewTransactionRequest{
		TransactionType: "withdrawal",
		Amount:          100,
	}

	AppError := request.Validate(101)

	if AppError != nil {
		t.Error("should not create err message when testing withdraw account")
	}
}

// test account amount as positif
func Test_err_if_account_amount_is_positif(t *testing.T) {
	request := NewTransactionRequest{
		TransactionType: "we don t care here",
		Amount:          -12,
	}

	AppError := request.Validate(20)

	if AppError.Message != "Amount can not be negatif" {
		t.Error("error testing Account negatif amount")
	}
}
