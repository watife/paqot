package customer

// Service defines the courier services
type Service interface {
	CreateCustomer(c *Customer) (*Customer, error)
	FindCustomerByID(id ID) (*Customer, error)
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

func (s *service) CreateCustomer(c *Customer) (*Customer, error) {
	cor, err := NewCustomer(c)

	if err != nil {
		return nil, err
	}

	return s.repo.Create(cor)
}

func (s *service) FindCustomerByID(id ID) (*Customer, error) {
	return s.repo.FindByID(id)
}
