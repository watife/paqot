package tests

import (
	"github.com/fakorede-bolu/deliva/api/courier"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCourier(t *testing.T) {
	c := NewFixtureCourier()
	u, err := courier.NewCourier(c)
	assert.Nil(t, err)
	assert.Equal(t, u.FirstName, "Boluwatife")
	assert.NotNil(t, u.ID)
}