package jobs

type Repository interface {
	Create(o *Jobs) (*Jobs, error)
	FindCustomerJob(ID, customerID ID) (*Jobs, error)
	FindJobByID(ID ID) (*Jobs, error)
	FindAllCustomerJobs(customerID ID) ([]*Jobs, error)
	CancelJob(ID, customerID ID) (string, error)
	VerifyJob(ID, customerID ID) (string, error)
	AssignCourierToJob(ID, courierID ID) (bool, error)
}
