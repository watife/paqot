package tests

import (
	"github.com/fakorede-bolu/deliva/api/courier"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func NewFixtureCourier() *courier.Courier {
	return &courier.Courier{
		ID:             courier.NewID(),
		FirstName:      "Ozzy",
		LastName:       "Osbourne",
		CreatedAt:      time.Now(),
		CountryCode:    "+234",
		DOB:            "1996-03-17",
		Address:        "ibadan, Nigeria",
		PhoneNumber:    8089333186,
		AffiliatedWith: "gig logistics",
	}
}

func TestService_Create(t *testing.T) {
	repo := newInmem()
	m := courier.NewService(repo)
	c := NewFixtureCourier()
	_, err := m.CreateCourier(c)
	assert.Nil(t, err)
	assert.False(t, c.CreatedAt.IsZero())
	assert.True(t, c.UpdatedAt.IsZero())
}
