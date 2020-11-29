package jobs

import "deliva/api/entities"

//inmem in memory repo
type inmem struct {
	m map[entities.ID]*entities.Jobs
}

func (i inmem) Create(o *entities.Jobs) (*entities.Jobs, error) {
	i.m[o.ID] = o
	return o, nil
}

func (i inmem) FindCustomerJob(customerID entities.ID) (*entities.Jobs, error) {
	panic("implement me")
}

func (i inmem) FindJobByID(ID entities.ID) (*entities.Jobs, error) {
	if i.m[ID] == nil {
		return nil, ErrNotFound
	}
	return i.m[ID], nil
}

func (i inmem) FindAllCustomerJobs(customerID entities.ID) ([]*entities.Jobs, error) {
	panic("implement me")
}

func (i inmem) CancelJob(ID, customerID entities.ID) (string, error) {
	panic("implement me")
}

func (i inmem) VerifyJob(ID, customerID entities.ID) (string, error) {
	panic("implement me")
}

func (i inmem) AssignCourierToJob(ID, courierID entities.ID) (bool, error) {
	if i.m[ID] == nil {
		return false, ErrNotFound
	}
	return true, nil
}

//NewInmem create new repository
func NewInmem() *inmem {
	var m = map[entities.ID]*entities.Jobs{}
	return &inmem{
		m: m,
	}
}