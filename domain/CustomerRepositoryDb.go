package domain

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/my-org/my-package/errs"
	"os"
	"time"
)

type CustomerRepositoryDB struct {
	client *sql.DB
}

func (d CustomerRepositoryDB) FindAll(filters map[string]string) ([]Customer, *errs.AppError) {

	findAllSql := "select id , name , city , zipcode , age , status from customers "

	var extras string

	switch filters["status"] {
	case "active":
		extras = extras + " where status = 1"
	case "inactive":
		extras = extras + " where status = 0"
	}

	findAllSql += extras

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	customers := make([]Customer, 0)

	for rows.Next() {

		var c Customer

		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Age, &c.Status)

		if err != nil {
			return nil, errs.NewUnexpectedError(err.Error())
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (d CustomerRepositoryDB) FindById(id int) (*Customer, *errs.AppError) {

	customerSql := "select id , name , city , zipcode , age , status from customers where id = ?"

	row := d.client.QueryRow(customerSql, id)

	var c Customer

	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.Age, &c.Status)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else {
			return nil, errs.NewUnexpectedError("unexpected Database Error")
		}
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDB {

	var dbUser string = os.Getenv("DB_USER")
	var dbPassword string = os.Getenv("DB_PASS")
	var dbAddr string = os.Getenv("DB_ADDR")
	var dbPort string = os.Getenv("DB_PORT")
	var dbName string = os.Getenv("DB_NAME")

	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)

	client, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetConnMaxIdleTime(10)

	return CustomerRepositoryDB{client: client}
}
