package customer

//inmem in memory repo
type inmem struct {
	m map[ID]*Customer
}

func (i inmem) Create(c *Customer) (*Customer, error) {
	i.m[c.ID] = c
	return c, nil
}

func (i inmem) FindByID(id ID) (*Customer, error) {
	if i.m[id] == nil {
		return nil, ErrCustomerNotFound
	}
	return i.m[id], nil
}

func (i inmem) FindByPhoneNumber(phoneNumber int, countryCode string) (*Customer, error) {
	panic("implement me")
}

//newInmem create new repository
func NewInmem() *inmem {
	var m = map[ID]*Customer{}
	return &inmem{
		m: m,
	}
}
