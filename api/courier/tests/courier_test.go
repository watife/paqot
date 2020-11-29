package tests

import (
	"deliva/api/entities"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCourier(t *testing.T) {
	c := NewFixtureCourier()
	u, err := entities.NewCourier(c)
	assert.Nil(t, err)
	assert.Equal(t, u.FirstName, "Boluwatife")
	assert.NotNil(t, u.ID)
}