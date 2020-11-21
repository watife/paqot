package customer

import "errors"

var  (
	ErrCustomerExists = errors.New("This customer already exist")
	ErrCustomerNotFound = errors.New("This customer does not exist")
)