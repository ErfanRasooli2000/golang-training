package service

import (
	"github.com/my-org/my-package/domain"
	"github.com/my-org/my-package/dto"
	"github.com/my-org/my-package/errs"
	"time"
)

type accountService interface {
	saveAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (receiver DefaultAccountService) saveAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {

	a := domain.Account{
		Id:          0,
		CustomerId:  request.CustomerId,
		OpeningDate: time.Now().Format("2006-03-02 23:23:23"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      true,
	}

	account, err := receiver.repo.Save(a)

	if err != nil {
		return nil, err
	}

	response := account.ToNewAccountResponseDto()

	return response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {

	return DefaultAccountService{repo: repo}

}
