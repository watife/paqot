package tests

import (
	"deliva/api/courier"
	"deliva/api/entities"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func NewFixtureCourier() *entities.Courier {
	return &entities.Courier{
		ID:             entities.NewID(),
		FirstName:      "Boluwatife",
		LastName:       "Fakorede",
		CreatedAt:      time.Now(),
		CountryCode:    "+234",
		DOB:            "1996-03-17",
		Address:        "ibadan, Nigeria",
		PhoneNumber:    8089333186,
		AffiliatedWith: "gig logistics",
	}
}

func TestService_Create(t *testing.T) {
	repo := courier.NewInmem()
	m := courier.NewService(repo)
	c := NewFixtureCourier()
	_, err := m.CreateCourier(c)
	assert.Nil(t, err)
	assert.False(t, c.CreatedAt.IsZero())
	assert.True(t, c.UpdatedAt.IsZero())
}
