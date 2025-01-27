package domain

import (
	"database/sql"
	"errors"
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

func (d CustomerRepositoryDB) FindById(id int) (*Customer, error) {

	customerSql := "select id , name , city , zipcode , age , status from customers where id = ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Age, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Customer Not Found")
		} else {
			return nil, errors.New("Unexpected Database Error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDB {
	client, err := sql.Open("mysql", "rosot:87438743@/banking")

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetConnMaxIdleTime(10)

	return CustomerRepositoryDB{client: client}
}
