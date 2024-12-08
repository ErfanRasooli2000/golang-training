package domain

import (
	"github.com/my-org/my-package/dto"
	"github.com/my-org/my-package/errs"
)

type Customer struct {
	Id          int `db:"customer_id"`
	Name        string
	DateOfBirth string `db:"date_of_birth"`
	City        string
	Zipcode     string
	Status      bool
}

type CustomerRepository interface {
	GetAll(filters map[string]string) ([]Customer, *errs.AppError)
	Find(id int) (Customer, *errs.AppError)
}

func (c Customer) ConvertCustomerToDto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		DateOfBirth: c.DateOfBirth,
		City:        c.City,
		Zipcode:     c.Zipcode,
		Status:      convertStatusToText(c.Status),
	}

}

func convertStatusToText(status bool) string {

	if status {
		return "active"
	}

	return "inactive"
}
