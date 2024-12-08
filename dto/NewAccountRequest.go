package dto

type NewAccountRequest struct {
	CustomerId  int     `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}
