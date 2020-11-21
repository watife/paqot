package customer

import (
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

func (r *Pg) Create(c *Customer) (*Customer, error) {
	cor, err := r.FindByPhoneNumber(c.PhoneNumber, c.CountryCode)

	if cor != nil {
		return nil, ErrCustomerExists
	}
	err = r.db.Create(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *Pg) FindByID(id ID) (*Customer, error) {
	cu := &Customer{}
	if err := r.db.Where("ID = ?", id ).Take(&cu).Error; err != nil {
		return nil, ErrCustomerNotFound
	}
	return cu, nil
}

func (r *Pg) FindByPhoneNumber(phoneNumber int, countryCode string) (*Customer, error) {
	cd := &Customer{}
	if err := r.db.Where("phone_number = ? AND country_code = ?",phoneNumber,countryCode).Take(&cd).Error; err != nil {
		return nil, ErrCustomerNotFound
	}
	return cd, nil
}


