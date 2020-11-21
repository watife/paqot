package courier

type Repository interface {
	Create(c *Courier) (*Courier, error)
	FindByID(id ID) (*Courier, error)
	FindByPhoneNumber(phoneNumber int, countryCode string) (*Courier, error)
	FindAll() ([]*Courier, error)
}