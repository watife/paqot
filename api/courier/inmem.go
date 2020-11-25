package courier

//inmem in memory repo
type inmem struct {
	m map[ID]*Courier
}

//newInmem create new repository
func NewInmem() *inmem {
	var m = map[ID]*Courier{}
	return &inmem{
		m: m,
	}
}

//Create a Courier
func (r *inmem) Create(c *Courier) (*Courier, error) {
	r.m[c.ID] = c
	return c, nil
}

func (r *inmem) FindByID(ID ID) (*Courier, error) {
	if r.m[ID] == nil {
		return nil, ErrCourierNotFound
	}
	return r.m[ID], nil
}

func (r *inmem) FindByPhoneNumber(phoneNumber int, countryCode string) (*Courier, error) {
	panic("implement me")
}

func (r *inmem) FindAll() ([]*Courier, error) {
	panic("implement me")
}
