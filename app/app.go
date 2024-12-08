package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/my-org/my-package/domain"
	"github.com/my-org/my-package/service"
	"log"
	"net/http"
	"os"
	"time"
)

func Start() {

	sanityCheck()

	router := mux.NewRouter()

	databaseConnectionClient := connectDb()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb(databaseConnectionClient))}

	router.HandleFunc("/customers", ch.Customers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.ShowCustomer).Methods(http.MethodGet)

	hostAddress := fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))

	err := http.ListenAndServe(hostAddress, router)

	if err != nil {
		log.Fatal(err)
	}
}

func sanityCheck() {

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	if address == "" || port == "" {

		panic("address or port have not been set")

	}

	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")

	if dbUser == "" || dbPass == "" || dbName == "" {

		panic("database credentials have not been set")

	}
}

func connectDb() *sqlx.DB {

	databaseCredentials := fmt.Sprintf("%s:%s@/%s", os.Getenv("DATABASE_USER"), os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"))

	db, err := sqlx.Open("mysql", databaseCredentials)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
