package jobs

import (
	"github.com/fakorede-bolu/deliva/api/courier"
	"github.com/fakorede-bolu/deliva/api/customer"
)

// Service defines the courier services
type Service interface {
	CreateJob(o *Jobs) (*Jobs, error)
	AssignCourierToJob(ID, courierID ID) (bool, error)
	FindJobByID(ID ID) (*Jobs, error)
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

func (s *service) CreateJob(o *Jobs) (*Jobs, error) {
	ord, err := NewJob(o)

	if err != nil {
		return nil, err
	}

	_, err = s.customerService.FindCustomerByID(ord.CustomerID)

	if err != nil {
		return nil, err
	}

	return s.repo.Create(ord)
}

func (s *service) AssignCourierToJob(ID, courierID ID) (bool, error) {
	_, err := s.courierService.FindByID(courierID)

	if err != nil {
		return false, err
	}

	ord, err := s.FindJobByID(ID)

	if err != nil {
		return false, err
	}

	if ord.Status != false {
		return false, ErrJobStatus
	}

	return s.repo.AssignCourierToJob(ID, courierID)
}

func  (s *service) FindJobByID(ID ID) (*Jobs, error){
	return s.repo.FindJobByID(ID)
}


