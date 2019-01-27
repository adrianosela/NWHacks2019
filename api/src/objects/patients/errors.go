package patients

import "errors"

var (
	//ErrPatientNotFound will be returned when a patient ID does not exist
	ErrPatientNotFound = errors.New("patient not found")
)
