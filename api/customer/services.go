package customer

import "deliva/api/entities"

// Service defines the courier services
type Service interface {
	CreateCustomer(c *entities.Customer) (*entities.Customer, error)
	FindCustomerByID(id entities.ID) (*entities.Customer, error)
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

func (s *service) CreateCustomer(c *entities.Customer) (*entities.Customer, error) {
	cor, err := entities.NewCustomer(c)

	if err != nil {
		return nil, err
	}

	return s.repo.Create(cor)
}

func (s *service) FindCustomerByID(id entities.ID) (*entities.Customer, error) {
	return s.repo.FindByID(id)
}
