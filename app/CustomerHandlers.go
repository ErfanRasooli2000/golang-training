package app

import (
	"encoding/json"
	"github.com/my-org/my-package/service"
	"net/http"
)

type Customer struct {
	Id      int `json:"id"`
	Name    int `json:"name"`
	City    int `json:"city"`
	Zipcode int `json:"zip_code"`
	Age     int `json:"age"`
	Status  int `json:"is_active"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(writer http.ResponseWriter, request *http.Request) {

	customers, _ := ch.service.GetAllCustomers()

	json.NewEncoder(writer).Encode(customers)

}
