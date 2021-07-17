package dto

import (
// "github.com/djedjethai/bankingSqlx/domain"
)

type NewTransactionRequest struct {
	AccountId       string
	Amount          float64
	TransactionType string
}

// func (d NewTransactionRequest) fromDtoToTransaction() domain.Transaction {
// 	return domain.Transaction{
// 		TransactionId:   "",
// 		AccountId:       d.AccountId,
// 		Amount:          d.Amount,
// 		TransactionType: d.TransactionType,
// 		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
// 	}
// }
