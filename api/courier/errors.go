package courier

import "errors"

var  (
	ErrCourierExists = errors.New("This courier already exist")
	ErrCourierNotFound = errors.New("This courier does not exist")
	ErrUpdate = errors.New("Could not update courier")
)