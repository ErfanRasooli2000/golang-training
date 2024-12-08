package domain

import (
	"github.com/my-org/my-package/dto"
	"github.com/my-org/my-package/errs"
)

type Account struct {
	Id          int
	CustomerId  int
	OpeningDate string
	AccountType string
	Amount      float64
	Status      bool
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {

	return dto.NewAccountResponse{
		AccountId: a.Id,
	}

}
