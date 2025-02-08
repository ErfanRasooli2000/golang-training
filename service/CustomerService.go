package service

import (
	"github.com/my-org/my-package/domain"
	"github.com/my-org/my-package/dto"
	"github.com/my-org/my-package/errs"
)

type CustomerService interface {
	GetAllCustomers(map[string]string) ([]domain.Customer, *errs.AppError)
	GetById(int) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(filters map[string]string) ([]domain.Customer, *errs.AppError) {
	return s.repo.FindAll(filters)
}

func (s DefaultCustomerService) GetById(id int) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)

	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repo}
}
