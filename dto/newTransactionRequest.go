package dto

import "github.com/djedjethai/bankingSqlx/domain"

type NewTransactionRequest struct {
	AccountId       string
	Amount          float64
	TransactionType string
}

func (d NewTransactionRequest) fromDtoToTransaction() domain.Transaction {
	// a finnir......
}
