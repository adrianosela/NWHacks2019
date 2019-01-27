package store

import (
	"github.com/adrianosela/NWHacks2019/api/src/objects/doctors"
	"github.com/adrianosela/NWHacks2019/api/src/objects/patients"
	"github.com/adrianosela/NWHacks2019/api/src/objects/prescriptions"
)

// DB represents the set of functions a datastore must implement in order
// to satisfy the API's requirements
type DB interface {
	// creators
	PutPatient(*patients.Patient) error
	PutDoctor(*doctors.Doctor) error
	PutPrescription(*prescriptions.Prescription) error
	// modifiers
	UpdatePatient(*patients.Patient) error
	UpdateDoctor(*doctors.Doctor) error
	UpdatePrescription(*prescriptions.Prescription) error
	// getters
	GetPatient(string) (*patients.Patient, error)
	GetDoctor(string) (*doctors.Doctor, error)
	GetPrescription(string) (*prescriptions.Prescription, error)
}
