package jobs

import "errors"

var  (
	ErrJobFailed = errors.New("Job creation failed")
	ErrAssigningCourier = errors.New("Courier could not be assigned to order")
	ErrNotFound = errors.New("Job not found")
	ErrJobStatus = errors.New("Job already has an active courier assigned, cancel to re-assign")
	ErrTooManyJobs = errors.New("You currently have a pending job")
	ErrCourierAvailable = errors.New("Your assigned courier is not available")
)
