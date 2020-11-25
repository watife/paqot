package tests

import (
	"github.com/fakorede-bolu/deliva/api/jobs"
	"github.com/fakorede-bolu/deliva/pkg/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func NewFakeJob() *jobs.Jobs  {
	return &jobs.Jobs{
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

func NewFakeJobCourier() *jobs.JobCourier {
	return &jobs.JobCourier{
		ID: jobs.NewID(),
		CourierID: jobs.NewID(),
	}
}

func TestNewJob(t *testing.T) {
	c := NewFakeJob()
	u, err := jobs.NewJob(c)
	assert.Nil(t, err)
	assert.NotNil(t, u.ID)

	if err != nil {
		assert.Equal(t,jobs.ErrJobFailed, err)
	}
}