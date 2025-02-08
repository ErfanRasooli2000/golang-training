package dto

type CustomerResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"full_name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Age     int    `json:"age"`
	Status  string `json:"status"`
}
