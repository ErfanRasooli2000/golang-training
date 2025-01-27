package domain

type Customer struct {
	Id      int
	Name    string
	City    string
	Zipcode string
	Age     int
	Status  bool
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
