package domain

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/my-org/my-package/errs"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) GetAll(filters map[string]string) ([]Customer, *errs.AppError) {

	findAllSql := "SELECT customer_id , name , city , zipcode , date_of_birth , status from customers"

	if status, ok := filters["status"]; ok {

		if status == "active" {

			findAllSql += " where status = 1"
		} else {
			findAllSql += " where status = 0"
		}
	}

	customers := make([]Customer, 0)
	err := d.client.Select(&customers, findAllSql)

	if err != nil {
		return nil, errs.ServerError(err.Error())
	}

	return customers, nil
}

func (d CustomerRepositoryDb) Find(id int) (Customer, *errs.AppError) {

	findQuerySql := `SELECT customer_id , name , city , zipcode , date_of_birth , status from customers where customer_id=?;`

	var c Customer

	err := d.client.Get(&c, findQuerySql, id)

	if err == nil {

		return c, nil

	} else if errors.Is(err, sql.ErrNoRows) {

		return c, errs.NotFoundHttpError("Not Found")

	} else {

		return c, errs.ServerError(err.Error())
	}
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client: dbClient}
}
