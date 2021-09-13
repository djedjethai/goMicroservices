package service

import (
	realDomain "github.com/djedjethai/bankingSqlx/domain"
	"github.com/djedjethai/bankingSqlx/dto"
	"github.com/djedjethai/bankingSqlx/errs"
	"github.com/djedjethai/bankingSqlx/mocks/domain"
	"github.com/golang/mock/gomock"
	"testing"
	"time"
)

// const dbTSLayout = "2006-01-02 15:04:05"
var mockRepo *domain.MockAccountRepository
var service AccountService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockRepo = domain.NewMockAccountRepository(ctrl)
	service = NewAccountService(mockRepo)

	return func() {
		service = nil
		defer ctrl.Finish()
	}
}

func Test_should_return_an_err_when_request_is_not_validated(t *testing.T) {
	// Arrange
	request := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      0,
	}

	service := NewAccountService(nil)

	// Act
	_, appErr := service.NewAccount(request)

	// Assert
	if appErr == nil {
		t.Error("failed while testing the new account validation")
	}
}

func Test_should_return_err_from_the_server_side_if_new_account_can_not_be_created(t *testing.T) {
	// Arrange
	tearDown := setup(t)
	defer tearDown()
	// ctrl := gomock.NewController(t)
	// defer ctrl.Finish()
	// mockRepo := domain.NewMockAccountRepository(ctrl)

	// create NewAccoutService() from our service
	// service := NewAccountService(mockRepo)
	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      5500,
	}
	account := realDomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	// THIS IS THE MOCK, THAT DEFINE THE EXPECTED MOCKED FUNC RESPONSE
	mockRepo.EXPECT().Save(account).Return(nil, errs.NewInternalServerError("db error from test"))

	// Act
	_, appErr := service.NewAccount(req)

	// Assert
	if appErr == nil {
		t.Error("Test failed while validating error for new account")
	}
}

func Test_should_return_new_account_if_new_account_has_been_created(t *testing.T) {
	// Arrange
	tearDown := setup(t)
	defer tearDown()

	req := dto.NewAccountRequest{
		CustomerId:  "100",
		AccountType: "saving",
		Amount:      5500,
	}
	account := realDomain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}
	accountWithId := account
	accountWithId.AccountId = "201"

	mockRepo.EXPECT().Save(account).Return(&accountWithId, nil)

	// Act
	newAccount, errApp := service.NewAccount(req)

	// Assert
	if errApp != nil {
		t.Error("Error should be null when saving match account")
	}

	if newAccount.AccountId != accountWithId.AccountId {
		t.Error("while testing 'should return newAccount' returned account should be same the created one")
	}
}
