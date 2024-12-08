package service

import (
	"github.com/my-org/my-package/domain"
	"github.com/my-org/my-package/dto"
	"github.com/my-org/my-package/errs"
)

type CustomerService interface {
	GetAllCustomers(filters map[string]string) ([]dto.CustomerResponse, *errs.AppError)
	FindCustomer(id int) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers(filters map[string]string) ([]dto.CustomerResponse, *errs.AppError) {

	customers, _ := service.repo.GetAll(filters)

	var response []dto.CustomerResponse = make([]dto.CustomerResponse, 0)

	for _, c := range customers {

		response = append(response, c.ConvertCustomerToDto())
	}

	return response, nil
}

func (service DefaultCustomerService) FindCustomer(id int) (*dto.CustomerResponse, *errs.AppError) {

	customer, err := service.repo.Find(id)

	if err != nil {
		return nil, err
	}

	response := customer.ConvertCustomerToDto()

	return &response, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {

	return DefaultCustomerService{repo: repo}
}
