package tests

import (
	"deliva/api/courier"
	"deliva/api/customer"
	"deliva/api/entities"
	"deliva/api/jobs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Create(t *testing.T) {
	repo := jobs.NewInmem()
	courierService := courier.NewService(courier.NewInmem())
	customerService := customer.NewService(customer.NewInmem())
	m := jobs.NewService(repo, courierService, customerService)
	j := NewFakeJob()

	t.Run("customer not found", func(t *testing.T) {
		c := &entities.Customer{
			ID: entities.NewID(),
		}

		_, err := customerService.FindCustomerByID(c.ID)
		_, err = m.CreateJob(j)
		assert.Equal(t, customer.ErrCustomerNotFound, err)
	})

	t.Run("success", func(t *testing.T) {
		c := &entities.Customer{
			ID: entities.NewID(),
		}

		cus, err := customerService.CreateCustomer(c)
		assert.Nil(t, err)

		_, err = customerService.FindCustomerByID(cus.ID)

		assert.Nil(t, err)

		if err != nil {
			assert.Equal(t, customer.ErrCustomerNotFound, err)
		}

		j.CustomerID = cus.ID
		_, err = m.CreateJob(j)

		if err != nil {
			assert.Equal(t,jobs.ErrJobFailed, err)
		}
		assert.Nil(t, err)
	})
}

func TestService_FindJobByID(t *testing.T)  {
	repo := jobs.NewInmem()
	courierService := courier.NewService(courier.NewInmem())
	customerService := customer.NewService(customer.NewInmem())
	m := jobs.NewService(repo, courierService, customerService)

	t.Run("job not found", func(t *testing.T) {
		j := &entities.Jobs{
			ID: entities.NewID(),
		}
		_, err := m.FindJobByID(j.ID)
		assert.Equal(t, jobs.ErrNotFound, err)
	})

	t.Run("success", func(t *testing.T) {
		j := NewFakeJob()
		c := &entities.Customer{
			ID: entities.NewID(),
		}
		cus, err := customerService.CreateCustomer(c)
		assert.Nil(t, err)

		j.CustomerID = cus.ID

		jo, err := m.CreateJob(j)
		assert.Nil(t, err)

		_, err = m.FindJobByID(jo.ID)

		assert.Nil(t, err)
	})


}

func TestService_AssignCourierToJob(t *testing.T)  {
	repo := jobs.NewInmem()
	courierService := courier.NewService(courier.NewInmem())
	customerService := customer.NewService(customer.NewInmem())
	m := jobs.NewService(repo, courierService, customerService)
	j := NewFakeJobCourier()
	jo := NewFakeJob()

	t.Run("job not found", func(t *testing.T) {
		c := &entities.Courier{
			ID: entities.NewID(),
		}
		_, err := m.AssignCourierToJob(jo.ID, c.ID)
		assert.Equal(t, courier.ErrCourierNotFound, err)
	})

	t.Run("success", func(t *testing.T) {
		c := &entities.Courier{
			ID: entities.NewID(),
		}

		cus := &entities.Customer{
			ID: entities.NewID(),
		}

		cust, err := customerService.CreateCustomer(cus)
		assert.Nil(t, err)

		jo.CustomerID = cust.ID

		newJob, err := m.CreateJob(jo)
		assert.Nil(t, err)

		cou, err := courierService.CreateCourier(c)

		assert.Nil(t, err)

		j.CourierID = cou.ID
		j.ID = newJob.ID

		_, err = m.AssignCourierToJob(j.ID, j.CourierID)

		assert.Nil(t, err)

	})
}
