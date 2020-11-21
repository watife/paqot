package jobs

import (
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

func (r *Pg) Create(o *Jobs) (*Jobs, error) {
	err := r.db.Create(&o).Error

	if err != nil {
		return nil, ErrJobFailed
	}

	return o, nil
}

func (r *Pg) FindCustomerJob(ID, customerID ID) (*Jobs, error) {
	panic("implement me")
}

func (r *Pg) FindJobByID(ID ID) (*Jobs, error) {
	ord := &Jobs{}
	if err := r.db.Where("ID = ?", ID ).Take(&ord).Error; err != nil {
		return nil, ErrNotFound
	}
	return ord, nil
}

func (r *Pg) FindAllCustomerJobs(customerID ID) ([]*Jobs, error) {
	panic("implement me")
}

func (r *Pg) CancelJob(ID, customerID ID) (string, error) {
	panic("implement me")
}

func (r *Pg) VerifyJob(ID, customerID ID) (string, error) {
	panic("implement me")
}

func (r *Pg) AssignCourierToJob(ID, courierID ID) (bool, error) {
	o := &Jobs{}
	r.db.First(o)

	o.Status = true
	o.CourierID = courierID
	err := r.db.Save(o).Error

	if err != nil {
		return false, ErrAssigningCourier
	}

	return true, nil
}
