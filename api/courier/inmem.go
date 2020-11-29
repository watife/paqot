package courier

import "deliva/api/entities"

//inmem in memory repo
type inmem struct {
	m map[entities.ID]*entities.Courier
}

func (r *inmem) AvailabilityStatus(ID entities.ID) (bool, error) {
	panic("implement me")
}

//newInmem create new repository
func NewInmem() *inmem {
	var m = map[entities.ID]*entities.Courier{}
	return &inmem{
		m: m,
	}
}

//Create a Courier
func (r *inmem) Create(c *entities.Courier) (*entities.Courier, error) {
	r.m[c.ID] = c
	return c, nil
}

func (r *inmem) FindByID(ID entities.ID) (*entities.Courier, error) {
	if r.m[ID] == nil {
		return nil, ErrCourierNotFound
	}
	return r.m[ID], nil
}

func (r *inmem) FindByPhoneNumber(phoneNumber int, countryCode string) (*entities.Courier, error) {
	panic("implement me")
}

func (r *inmem) FindAll() ([]*entities.Courier, error) {
	panic("implement me")
}
