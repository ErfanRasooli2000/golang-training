package app

import (
	"github.com/gorilla/mux"
	"github.com/my-org/my-package/service"
	"net/http"
	"strconv"
)

type Customer struct {
	Id          int    `json:"customer_id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date-of-birth"`
	City        string `json:"city"`
	Zipcode     string `json:"zip_code"`
	Status      string `json:"status"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) Customers(writer http.ResponseWriter, request *http.Request) {

	var filters map[string]string
	filters = make(map[string]string)

	status := request.URL.Query().Get("status")

	if status == "active" || status == "inactive" {

		filters["status"] = status
	}

	customers, err := ch.service.GetAllCustomers(filters)

	if err == nil {
		writeResponse(writer, http.StatusOK, customers)

	} else {
		writeResponse(writer, err.Code, err)
	}
}

func (ch *CustomerHandlers) ShowCustomer(writer http.ResponseWriter, request *http.Request) {

	inputs := mux.Vars(request)

	id, _ := strconv.Atoi(inputs["id"])

	customer, err := ch.service.FindCustomer(id)

	if err == nil {
		writeResponse(writer, http.StatusOK, customer)

	} else {
		writeResponse(writer, err.Code, err)
	}
}
