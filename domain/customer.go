package domain

import (
	"github.com/my-org/my-package/dto"
	"github.com/my-org/my-package/errs"
)

type Customer struct {
	Id      int
	Name    string
	City    string
	Zipcode string
	Age     int
	Status  bool
}

func (c Customer) statusAsText() string {

	if c.Status {
		return "active"
	}

	return "inactive"
}

func (c Customer) ToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:      c.Id,
		Name:    c.Name,
		City:    c.City,
		Zipcode: c.Zipcode,
		Age:     c.Age,
		Status:  c.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(map[string]string) ([]Customer, *errs.AppError)
	FindById(int) (*Customer, *errs.AppError)
}
