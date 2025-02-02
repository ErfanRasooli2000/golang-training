package app

import (
	"github.com/gorilla/mux"
	"github.com/my-org/my-package/domain"
	"github.com/my-org/my-package/service"
	"net/http"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	http.ListenAndServe(":8080", router)

}
