package jobs

import (
	"time"
)

type Jobs struct {
	ID             	ID
	CustomerID     	ID `json:"customerID" db:"customer_id" validate:"required"`
	CourierID     	ID `json:"courierID" db:"courier_id"`
	Description    	string `json:"description" db:"description" validate:"required"`
	OrderImage    	string `json:"orderImage" db:"order_image"`
	DeliveryAddress string `json:"deliveryAddress" db:"delivery_address" validate:"required"`
	PickUpLat		float64 `json:"pickUpLat" db:"pickup_lat" validate:"required"`
	PickUpLong		float64 `json:"pickUpLong" db:"pickup_long" validate:"required"`
	DeliveryLat     float64 `json:"deliveryLat" db:"delivery_lat" validate:"required"`
	DeliveryLong    float64 `json:"deliveryLong" db:"delivery_long" validate:"required"`
	Price    	   	float64 `json:"price" db:"price" validate:"required"`
	CreatedAt      	time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt      	time.Time  `json:"updatedAt" db:"created_at"`
	Verified 	   	bool `json:"verified" db:"verified"`
	Status 	   		bool `json:"status" db:"status"`
}

type JobCourier struct {
	ID  ID `json:"id" db:"id" validate:"required"`
	CourierID ID `json:"courierId" db:"courier_id" validate:"required"`
}

func NewJob(j *Jobs) (*Jobs, error)  {
	j.ID = NewID()
	return j, nil
}

