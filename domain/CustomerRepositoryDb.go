package domain

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAllSql := "select id , name , city , zipcode , age , status from customers"

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		panic(err)
	}

	customers := make([]Customer, 0)

	for rows.Next() {

		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Age, &c.Status)

		if err != nil {
			panic(err)
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "root:87438743@/banking")

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetConnMaxIdleTime(10)

	return CustomerRepositoryDB{client: client}
}
