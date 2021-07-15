package dto

type NewTransactionResponse struct {
	Amount        float64 `json:"amount"`
	TransactionId string  `json:"transaction_id"`
}
