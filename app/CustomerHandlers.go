package app

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/my-org/my-package/service"
	"net/http"
	"strconv"
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
func (ch *CustomerHandlers) getCustomerById(writer http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	id, _ := strconv.Atoi(vars["id"])

	customer, err := ch.service.GetById(id)

	if err != nil {

		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, err.Error())

	} else {
		writer.Header().Add("Content-Type", "Application/json")
		json.NewEncoder(writer).Encode(customer)
	}

}
