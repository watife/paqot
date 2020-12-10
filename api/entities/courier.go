package entities

import (
	"time"
)

type Courier struct {
	ID             ID
	FirstName      string     `json:"firstName" db:"first_name" validate:"required"`
	LastName       string     `json:"lastName" db:"last_name" validate:"required"`
	CountryCode    string     `json:"countryCode" db:"country_code" validate:"min=1,max=4,required"`
	PhoneNumber    int        `json:"phoneNumber" db:"phone_number" validate:"required"`
	Rating         float32    `json:"rating" db:"rating"`
	DOB            string     `json:"DOB" db:"DOB" validate:"required"`
	Guarantor      string     `json:"guarantor" db:"guarantor"`
	Address        string     `json:"address" db:"address" validate:"required"`
	ProfilePicture string     `json:"profilePicture" db:"profile_picture"`
	Identifier     string     `json:"identifier" db:"identifier"`
	AffiliatedWith string 	  `json:"affiliatedWith" db:"affiliated_with" validate:"-"` //signifies if associated with any company
	CreatedAt      time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt      time.Time  `json:"updatedAt" db:"created_at"`
	Verified 	   bool 	  `json:"verified" db:"verified"`
	Availability   bool       `json:"availability" db:"availability" gorm:"default:true"`
	Job 		   []Jobs
}

func NewCourier(c *Courier) (*Courier, error)  {
	c.ID = NewID()
	return c, nil

}
