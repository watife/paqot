package customer

import "deliva/api/entities"

//inmem in memory repo
type inmem struct {
	m map[entities.ID]*entities.Customer
}

func (i inmem) Create(c *entities.Customer) (*entities.Customer, error) {
	i.m[c.ID] = c
	return c, nil
}

func (i inmem) FindByID(id entities.ID) (*entities.Customer, error) {
	if i.m[id] == nil {
		return nil, ErrCustomerNotFound
	}
	return i.m[id], nil
}

func (i inmem) FindByPhoneNumber(phoneNumber int, countryCode string) (*entities.Customer, error) {
	panic("implement me")
}

//newInmem create new repository
func NewInmem() *inmem {
	var m = map[entities.ID]*entities.Customer{}
	return &inmem{
		m: m,
	}
}
