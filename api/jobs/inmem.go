package jobs

//inmem in memory repo
type inmem struct {
	m map[ID]*Jobs
}

func (i inmem) Create(o *Jobs) (*Jobs, error) {
	i.m[o.ID] = o
	return o, nil
}

func (i inmem) FindCustomerJob(ID, customerID ID) (*Jobs, error) {
	panic("implement me")
}

func (i inmem) FindJobByID(ID ID) (*Jobs, error) {
	if i.m[ID] == nil {
		return nil, ErrNotFound
	}
	return i.m[ID], nil
}

func (i inmem) FindAllCustomerJobs(customerID ID) ([]*Jobs, error) {
	panic("implement me")
}

func (i inmem) CancelJob(ID, customerID ID) (string, error) {
	panic("implement me")
}

func (i inmem) VerifyJob(ID, customerID ID) (string, error) {
	panic("implement me")
}

func (i inmem) AssignCourierToJob(ID, courierID ID) (bool, error) {
	if i.m[ID] == nil {
		return false, ErrNotFound
	}
	return true, nil
}

//NewInmem create new repository
func NewInmem() *inmem {
	var m = map[ID]*Jobs{}
	return &inmem{
		m: m,
	}
}