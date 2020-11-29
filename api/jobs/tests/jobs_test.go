package tests

import (
	"deliva/api/entities"
	"deliva/api/jobs"
	"deliva/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewFakeJob() *entities.Jobs {
	return &entities.Jobs{
		ID: helpers.NewID(),
		CustomerID: helpers.NewID(),
		Description: "I need someone to help deliver this product",
		DeliveryAddress: "3H, Courier services",
		PickUpLat: 7.400470,
		PickUpLong: 3.872350,
		DeliveryLat: 39.484940,
		DeliveryLong: -121.220467,
		Price: 44,
	}
}

func NewFakeJobCourier() *entities.JobCourier {
	return &entities.JobCourier{
		ID: entities.NewID(),
		CourierID: entities.NewID(),
	}
}

func TestNewJob(t *testing.T) {
	c := NewFakeJob()
	u, err := entities.NewJob(c)
	assert.Nil(t, err)
	assert.NotNil(t, u.ID)

	if err != nil {
		assert.Equal(t,jobs.ErrJobFailed, err)
	}
}