package store

import (
	"github.com/adrianosela/NWHacks2019/api/src/objects/doctors"
	"github.com/adrianosela/NWHacks2019/api/src/objects/patients"
	"github.com/adrianosela/NWHacks2019/api/src/objects/prescriptions"
)

// MockDB is an in-memory datastore for prototyping
type MockDB struct {
	Doctors      map[string]doctors.Doctor
	Patients     map[string]patients.Patient
	Prescription map[string]prescriptions.Prescription
}

// NewMockDB returns a mock database which implements the DB interface
func NewMockDB() *MockDB {
	return &MockDB{
		Doctors:      make(map[string]doctors.Doctor),
		Patients:     make(map[string]patients.Patient),
		Prescription: make(map[string]prescriptions.Prescription),
	}
}

// PutDoctor stores a doctor in the db
func (db *MockDB) PutDoctor(dr *doctors.Doctor) error {
	if _, exists := db.Doctors[dr.ID]; exists {
		return ErrItemExists
	}
	db.Doctors[dr.ID] = *dr
	return nil
}

// PutPatient stores a patient in the db
func (db *MockDB) PutPatient(pt *patients.Patient) error {
	if _, exists := db.Patients[pt.ID]; exists {
		return ErrItemExists
	}
	db.Patients[pt.ID] = *pt
	return nil
}

// PutPrescription stores a prescription in the db
func (db *MockDB) PutPrescription(pr *prescriptions.Prescription) error {
	if _, exists := db.Prescription[pr.ID]; exists {
		return ErrItemExists
	}
	db.Prescription[pr.ID] = *pr
	return nil
}

// UpdateDoctor updates a doctor in the db
func (db *MockDB) UpdateDoctor(dr *doctors.Doctor) error {
	db.Doctors[dr.ID] = *dr
	return nil
}

// UpdatePatient updates a doctor in the db
func (db *MockDB) UpdatePatient(pt *patients.Patient) error {
	db.Patients[pt.ID] = *pt
	return nil
}

// UpdatePrescription updates a prescription in the db
func (db *MockDB) UpdatePrescription(pr *prescriptions.Prescription) error {
	db.Prescription[pr.ID] = *pr
	return nil
}

// GetDoctor gets a doctor from the db
func (db *MockDB) GetDoctor(drID string) (*doctors.Doctor, error) {
	if val, ok := db.Doctors[drID]; ok {
		return &val, nil
	}
	return nil, ErrNotInStore
}

// GetPatient gets a patient from the db
func (db *MockDB) GetPatient(ptID string) (*patients.Patient, error) {
	if val, ok := db.Patients[ptID]; ok {
		return &val, nil
	}
	return nil, ErrNotInStore
}

// GetPrescription gets a prescription from the db
func (db *MockDB) GetPrescription(prID string) (*prescriptions.Prescription, error) {
	if val, ok := db.Prescription[prID]; ok {
		return &val, nil
	}
	return nil, ErrNotInStore
}
