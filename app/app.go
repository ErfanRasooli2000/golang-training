package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/my-org/my-package/domain"
	"github.com/my-org/my-package/service"
	"net/http"
	"os"
)

func sanityCheck() {

	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASS") == "" ||
		os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" {

		panic("env error")

	}
}

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	var address string = os.Getenv("SERVER_ADDRESS")
	var port string = os.Getenv("SERVER_PORT")

	var serve string = fmt.Sprintf("%s:%s", address, port)

	http.ListenAndServe(serve, router)

}
