package dto

import (
	"net/http"
	"testing"
)

func Test_amount_newAccountRequest(t *testing.T) {
	// Arrange
	request := NewAccountRequest{
		AccountType: "saving",
		Amount:      4000,
	}

	// Act
	AppErr := request.Validate()

	// Assert
	if AppErr.Message != "To open an account 5000$ is require" {
		t.Error("Wrong err message testing minimum amount while openNewAccount ")
	}
	if AppErr.Code != http.StatusUnprocessableEntity {
		t.Error("Wrong err message testing minimum amount while openNewAccount ")
	}
}

func Test_type_of_account(t *testing.T) {
	request := NewAccountRequest{
		AccountType: "wrong type",
		Amount:      5500,
	}

	AppErr := request.Validate()

	if AppErr.Message != "Account should be 'saving' or 'checking'" {
		t.Error("Wrong type of account when testing NewAccountRequest")
	}
}
