package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/my-org/my-package/Logger"
	"github.com/my-org/my-package/errs"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(account Account) (*Account, *errs.AppError) {

	sqlInsert := "INSERT INTO accounts (customer_id , opening_date , account_type , amount , status) values (?,?,?,?,?)"

	result, err := d.client.Exec(sqlInsert,
		account.CustomerId,
		account.OpeningDate,
		account.AccountType,
		account.Amount,
		account.Status,
	)

	if err != nil {
		Logger.Error(err.Error())
		return nil, errs.ServerError(err.Error())
	}

	id, _ := result.LastInsertId()

	account.Id = int(id)

	return &account, nil
}

func newAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: dbClient}
}
