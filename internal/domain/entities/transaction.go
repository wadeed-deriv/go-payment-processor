package entities

type Transaction struct {
	ClientID int     `json:"client_id"`
	Type     string  `json:"type"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
}
