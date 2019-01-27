package prescriptions

import "errors"

var (
	//ErrPrescriptionDoesNotExist will be returned when a prescription ID does not exist
	ErrPrescriptionDoesNotExist = errors.New("prescription does not exist")
)
