package store

import (
	"errors"

	"github.com/adrianosela/NWHacks2019/api/src/objects/doctors"
	"github.com/adrianosela/NWHacks2019/api/src/objects/patients"
	"github.com/adrianosela/NWHacks2019/api/src/objects/prescriptions"
)

var (
	// ErrItemExists will be returned when the developer attempts to overwite an item
	// with a put function instead of a modify (item already exists)
	ErrItemExists = errors.New("item already exists, use update function to change")
	// ErrNotInStore will be returned when the developer attempts retrieve an item not in the db
	ErrNotInStore = errors.New("the requested item does not exist in the db")
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
