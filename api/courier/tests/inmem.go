package tests

import "github.com/fakorede-bolu/deliva/api/courier"

//inmem in memory repo
type inmem struct {
	m map[courier.ID]*courier.Courier
}

//newInmem create new repository
func newInmem() *inmem {
	var m = map[courier.ID]*courier.Courier{}
	return &inmem{
		m: m,
	}
}

//Create a Courier
func (r *inmem) Create(c *courier.Courier) (*courier.Courier, error) {
	r.m[c.ID] = c
	return c, nil
}

func (r *inmem) FindByID(id string) (*courier.Courier, error) {
	panic("implement me")
}

func (r *inmem) FindByPhoneNumber(phoneNumber int, countryCode string) (*courier.Courier, error) {
	panic("implement me")
}

func (r *inmem) FindAll() ([]*courier.Courier, error) {
	panic("implement me")
}
