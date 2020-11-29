package courier

import "deliva/api/entities"

type Repository interface {
	Create(c *entities.Courier) (*entities.Courier, error)
	FindByID(id entities.ID) (*entities.Courier, error)
	FindByPhoneNumber(phoneNumber int, countryCode string) (*entities.Courier, error)
	FindAll() ([]*entities.Courier, error)
	AvailabilityStatus(ID entities.ID, status bool) (bool, error)
}