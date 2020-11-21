package customer

type Repository interface {
	Create(c *Customer) (*Customer, error)
	FindByID(id ID) (*Customer, error)
	FindByPhoneNumber(phoneNumber int, countryCode string) (*Customer, error)
}