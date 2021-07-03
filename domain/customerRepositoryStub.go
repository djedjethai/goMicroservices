package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Jeje", City: "paris", ZipCode: "75000", DateOfBirth: "2000-1-1", Status: "1"},
		{Id: "1002", Name: "Anna", City: "Bkk", ZipCode: "10110", DateOfBirth: "2002-2-2", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}
