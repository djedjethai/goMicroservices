// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/djedjethai/bankingSqlx/domain (interfaces: TransactionRepository)

// Package domain is a generated GoMock package.
package domain

import (
	reflect "reflect"

	domain "github.com/djedjethai/bankingSqlx/domain"
	dto "github.com/djedjethai/bankingSqlx/dto"
	errs "github.com/djedjethai/bankingSqlx/errs"
	gomock "github.com/golang/mock/gomock"
)

// MockTransactionRepository is a mock of TransactionRepository interface.
type MockTransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTransactionRepositoryMockRecorder
}

// MockTransactionRepositoryMockRecorder is the mock recorder for MockTransactionRepository.
type MockTransactionRepositoryMockRecorder struct {
	mock *MockTransactionRepository
}

// NewMockTransactionRepository creates a new mock instance.
func NewMockTransactionRepository(ctrl *gomock.Controller) *MockTransactionRepository {
	mock := &MockTransactionRepository{ctrl: ctrl}
	mock.recorder = &MockTransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTransactionRepository) EXPECT() *MockTransactionRepositoryMockRecorder {
	return m.recorder
}

// GetBalance mocks base method.
func (m *MockTransactionRepository) GetBalance(arg0 string) (float64, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", arg0)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockTransactionRepositoryMockRecorder) GetBalance(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockTransactionRepository)(nil).GetBalance), arg0)
}

// RunTransaction mocks base method.
func (m *MockTransactionRepository) RunTransaction(arg0 domain.Transaction, arg1 float64) (*dto.NewTransactionResponse, *errs.AppError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunTransaction", arg0, arg1)
	ret0, _ := ret[0].(*dto.NewTransactionResponse)
	ret1, _ := ret[1].(*errs.AppError)
	return ret0, ret1
}

// RunTransaction indicates an expected call of RunTransaction.
func (mr *MockTransactionRepositoryMockRecorder) RunTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunTransaction", reflect.TypeOf((*MockTransactionRepository)(nil).RunTransaction), arg0, arg1)
}
