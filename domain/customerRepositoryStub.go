package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {

	customers := []Customer{
		{Id: 1, Name: "Erfan", City: "Tehran", Zipcode: "11111111", Age: 25, Status: true},
		{Id: 2, Name: "Saeed", City: "Tabriz", Zipcode: "11111111", Age: 53, Status: false},
		{Id: 3, Name: "Farah", City: "Tabriz", Zipcode: "11111111", Age: 47, Status: false},
		{Id: 4, Name: "Elnaz", City: "Tehran", Zipcode: "11111111", Age: 13, Status: true},
	}

	return CustomerRepositoryStub{customers: customers}
}
