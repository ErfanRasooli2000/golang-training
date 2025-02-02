package domain

import "github.com/my-org/my-package/errs"

type Customer struct {
	Id      int
	Name    string
	City    string
	Zipcode string
	Age     int
	Status  bool
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(int) (*Customer, *errs.AppError)
}
