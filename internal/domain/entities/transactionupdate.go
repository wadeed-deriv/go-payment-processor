package entities

type TransactionUpdate struct {
	AccountID       string  `json:"accountid"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transactiontype"`
}
