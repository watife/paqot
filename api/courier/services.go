package courier

import "deliva/api/entities"

// Service defines the courier services
type Service interface {
	CreateCourier(c *entities.Courier) (*entities.Courier, error)
	FindByID(ID entities.ID) (*entities.Courier, error)
	ManageAvailabilityStatus(ID entities.ID, status bool) (bool, error)
}

//Service struct
type service struct {
	repo Repository
}

//NewService create new use case
func NewService(r Repository) *service {
	return &service{
		repo: r,
	}
}

func (s *service) CreateCourier(c *entities.Courier) (*entities.Courier, error) {
	cor, err := entities.NewCourier(c)

	if err != nil {
		return nil, err
	}

	return s.repo.Create(cor)
}

func (s *service) FindByID(ID entities.ID) (*entities.Courier, error)  {
	return s.repo.FindByID(ID)
}

func (s *service) ManageAvailabilityStatus(ID entities.ID, status bool) (bool, error)  {
	return s.repo.AvailabilityStatus(ID, status)
}