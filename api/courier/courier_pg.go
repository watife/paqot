package courier

import (
	"gorm.io/gorm"
)

type Pg struct {
	db   *gorm.DB
}

func NewCourierPg(db *gorm.DB) *Pg {
	return &Pg{
		db,
	}
}

func (r *Pg) Create(c *Courier) (*Courier, error) {
	cor, err := r.FindByPhoneNumber(c.PhoneNumber, c.CountryCode)

	if cor != nil {
		return nil, ErrCourierExists
	}
	err = r.db.Create(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *Pg) FindByID(id ID) (*Courier, error) {
	cd := &Courier{}
	if err := r.db.Where("ID = ?", id ).Take(&cd).Error; err != nil {
		return nil, ErrCourierNotFound
	}
	return cd, nil
}

func (r *Pg) FindAll() ([]*Courier, error) {
	panic("implement me")
}

func (r *Pg) FindByPhoneNumber(phoneNumber int, countryCode string) (*Courier, error) {
	cd := &Courier{}
	if err := r.db.Where("phone_number = ? AND country_code = ?",phoneNumber,countryCode).Take(&cd).Error; err != nil {
		return nil, err
	}
	return cd, nil
}


