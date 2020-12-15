package customer

import (
	"deliva/api/entities"
	"gorm.io/gorm"
)

type Pg struct {
	db   *gorm.DB
}

func NewCustomerPg(db *gorm.DB) *Pg {
	return &Pg{
		db,
	}
}

func (r *Pg) Create(c *entities.Customer) (*entities.Customer, error) {
	cor, _ := r.FindByPhoneNumber(c.PhoneNumber, c.CountryCode)

	if cor != nil {
		return nil, ErrCustomerExists
	}
	err := r.db.Create(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *Pg) FindByID(id entities.ID) (*entities.Customer, error) {
	cu := &entities.Customer{}
	if err := r.db.Preload("Jobs").Where("ID = ?", id ).Take(&cu).Error; err != nil {
		return nil, ErrCustomerNotFound
	}
	return cu, nil
}

func (r *Pg) FindByPhoneNumber(phoneNumber int, countryCode string) (*entities.Customer, error) {
	cd := &entities.Customer{}
	if err := r.db.Where("phone_number = ? AND country_code = ?",phoneNumber,countryCode).Take(&cd).Error; err != nil {
		return nil, ErrCustomerNotFound
	}
	return cd, nil
}


