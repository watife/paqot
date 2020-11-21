package courier

// Service defines the courier services
type Service interface {
	CreateCourier(c *Courier) (*Courier, error)
	FindByID(ID ID) (*Courier, error)
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

func (s *service) CreateCourier(c *Courier) (*Courier, error) {
	cor, err := NewCourier(c)

	if err != nil {
		return nil, err
	}

	return s.repo.Create(cor)
}

func (s *service) FindByID(ID ID) (*Courier, error)  {
	return s.repo.FindByID(ID)
}