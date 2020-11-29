package entities

import (
	"time"
)

type Customer struct {
	ID          ID
	FirstName   string    `json:"firstName" db:"first_name" validate:"required"`
	LastName    string    `json:"lastName" db:"last_name" validate:"required"`
	CountryCode string    `json:"countryCode" db:"country_code" validate:"min=1,max=4,required"`
	PhoneNumber int       `json:"phoneNumber" db:"phone_number" validate:"required"`
	Rating      float32   `json:"rating" db:"rating"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"created_at"`
	Verified    bool      `json:"verified" db:"verified"`
	Jobs        []Jobs
}

func NewCustomer(c *Customer) (*Customer, error)  {
	c.ID = NewID()
	return c, nil
}

