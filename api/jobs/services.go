package jobs

import (
	"deliva/api/courier"
	"deliva/api/customer"
	"deliva/api/entities"
)

// Service defines the courier services
type Service interface {
	CreateJob(o *entities.Jobs) (*entities.Jobs, error)
	AssignCourierToJob(ID, courierID entities.ID) (bool, error)
	FindJobByID(ID entities.ID) (*entities.Jobs, error)
	FindCustomerJob(customerID entities.ID) (*entities.Jobs, error)
}

//Service struct
type service struct {
	repo Repository
	courierService courier.Service
	customerService customer.Service
}

//NewService create new use case
func NewService(r Repository, c courier.Service, cu customer.Service) *service {
	return &service{
		repo: r,
		courierService: c,
		customerService: cu,
	}
}

func (s *service) CreateJob(o *entities.Jobs) (*entities.Jobs, error) {
	job, err := entities.NewJob(o)

	if err != nil {
		return nil, err
	}

	_, err = s.customerService.FindCustomerByID(job.CustomerID)

	if err != nil {
		return nil, err
	}

	jo, err := s.FindCustomerJob(job.CustomerID)

	if err != nil {
		return nil, err
	}

	if jo.Status != true {
		return nil, ErrTooManyJobs
	}

	return s.repo.Create(job)
}

func (s *service) AssignCourierToJob(ID, courierID entities.ID) (bool, error) {
	cou, err := s.courierService.FindByID(courierID)

	if err != nil {
		return false, err
	}

	if cou.Availability != true {
		return false, ErrCourierAvailable
	}

	jo, err := s.FindJobByID(ID)

	if err != nil {
		return false, err
	}

	if jo.Status != false {
		return false, ErrJobStatus
	}

	_, err = s.courierService.ManageAvailabilityStatus(courierID, false)

	if err != nil {
		return false, err
	}

	return s.repo.AssignCourierToJob(ID, courierID)
}

func  (s *service) FindJobByID(ID entities.ID) (*entities.Jobs, error){
	return s.repo.FindJobByID(ID)
}

func  (s *service) FindCustomerJob(customerID entities.ID) (*entities.Jobs, error){
	return s.repo.FindCustomerJob(customerID)
}



