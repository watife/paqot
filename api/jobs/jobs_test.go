package jobs

import (
	"deliva/api/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewFakeJob() *entities.Jobs {
	return &entities.Jobs{
		ID: entities.NewID(),
		CustomerID: entities.NewID(),
		Description: "I need someone to help deliver this product",
		DeliveryAddress: "3H, Courier services",
		PickUpLat: 7.400470,
		PickUpLong: 3.872350,
		DeliveryLat: 39.484940,
		DeliveryLong: -121.220467,
		Price: 44,
		Status: false,
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
		assert.Equal(t, ErrJobFailed, err)
	}
}