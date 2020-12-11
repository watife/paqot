package courier

import (
	"deliva/api/entities"
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

func (r *Pg) Create(c *entities.Courier) (*entities.Courier, error) {
	cor, _ := r.FindByPhoneNumber(c.PhoneNumber, c.CountryCode)

	if cor != nil {
		return nil, ErrCourierExists
	}
	err := r.db.Create(&c).Error

	if err != nil {
		return nil, err
	}

	return c, nil
}

func (r *Pg) FindByID(id entities.ID) (*entities.Courier, error) {
	cd := &entities.Courier{}
	if err := r.db.Where("ID = ?", id ).Take(&cd).Error; err != nil {
		return nil, ErrCourierNotFound
	}
	return cd, nil
}

func (r *Pg) FindAll() ([]*entities.Courier, error) {
	panic("implement me")
}

func (r *Pg) FindByPhoneNumber(phoneNumber int, countryCode string) (*entities.Courier, error) {
	cd := &entities.Courier{}
	if err := r.db.Where("phone_number = ? AND country_code = ?",phoneNumber,countryCode).Take(&cd).Error; err != nil {
		return nil, err
	}
	return cd, nil
}

func (r *Pg) AvailabilityStatus(ID entities.ID, status bool) (bool, error) {
	cd := &entities.Courier{}
	if err := r.db.Model(&cd).Where("id = ?", ID).Update("availability", status).Error; err != nil {
		return false, ErrUpdate
	}

	return true, nil
}
