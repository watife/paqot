package customer

import "deliva/api/entities"

type Repository interface {
	Create(c *entities.Customer) (*entities.Customer, error)
	FindByID(id entities.ID) (*entities.Customer, error)
	FindByPhoneNumber(phoneNumber int, countryCode string) (*entities.Customer, error)
}