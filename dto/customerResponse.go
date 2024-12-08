package dto

type CustomerResponse struct {
	Id          int    `json:"customer_id"`
	Name        string `json:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `json:"city"`
	Zipcode     string `json:"zip-code"`
	Status      string `json:"status"`
}
