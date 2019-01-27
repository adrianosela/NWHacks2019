package doctors

import "errors"

var (
	//ErrDoctorNotFound will be returned when a doctor ID does not exist
	ErrDoctorNotFound = errors.New("doctor not found")
)
