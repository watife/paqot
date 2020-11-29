package jobs

import (
	"deliva/api/entities"
	"gorm.io/gorm"
)

type Pg struct {
	db   *gorm.DB
}


func NewJobsPg(db *gorm.DB) *Pg {
	return &Pg{
		db,
	}
}

func (r *Pg) Create(o *entities.Jobs) (*entities.Jobs, error) {
	err := r.db.Create(&o).Error

	if err != nil {
		return nil, ErrJobFailed
	}

	return o, nil
}

func (r *Pg) FindCustomerJob(customerID entities.ID) (*entities.Jobs, error) {
	j := &entities.Jobs{}
	if err := r.db.Where("customer_id", customerID ).Last(&j).Error; err != nil {
		return nil, err
	}
	return j, nil
}

func (r *Pg) FindJobByID(ID entities.ID) (*entities.Jobs, error) {
	ord := &entities.Jobs{}
	if err := r.db.Where("ID = ?", ID ).Take(&ord).Error; err != nil {
		return nil, ErrNotFound
	}
	return ord, nil
}

func (r *Pg) FindAllCustomerJobs(customerID entities.ID) ([]*entities.Jobs, error) {
	panic("implement me")
}

func (r *Pg) CancelJob(ID, customerID entities.ID) (string, error) {
	panic("implement me")
}

func (r *Pg) VerifyJob(ID, customerID entities.ID) (string, error) {
	panic("implement me")
}

func (r *Pg) AssignCourierToJob(ID, courierID entities.ID) (bool, error) {
	o := &entities.Jobs{}
	r.db.First(o)

	o.Status = true
	o.CourierID = courierID
	err := r.db.Save(o).Error

	if err != nil {
		return false, ErrAssigningCourier
	}

	return true, nil
}
