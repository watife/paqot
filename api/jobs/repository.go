package jobs

import "deliva/api/entities"

type Repository interface {
	Create(o *entities.Jobs) (*entities.Jobs, error)
	FindCustomerLastJob(customerID entities.ID) (*entities.Jobs, error)
	FindJobByID(ID entities.ID) (*entities.Jobs, error)
	FindAllCustomerJobs(customerID entities.ID) ([]*entities.Jobs, error)
	CancelJob(ID, customerID entities.ID) (string, error)
	VerifyJob(ID, customerID entities.ID) (string, error)
	AssignCourierToJob(ID, courierID entities.ID) (bool, error)
	//CompleteJob(ID entities.ID) (bool, error) @todo also freeup the courier to take other jobs (ManageAvailabilityStatus service)
}
