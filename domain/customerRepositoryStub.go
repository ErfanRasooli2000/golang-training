package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) GetAll() ([]Customer, error) {

	return s.customers, nil

}

//func NewCustomerRepositoryStub() CustomerRepositoryStub {
//
//
//
//}
