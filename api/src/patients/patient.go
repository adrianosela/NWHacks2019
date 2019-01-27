package patients

import (
	"github.com/adrianosela/NWHacks2019/api/src/prescriptions"
	uuid "github.com/satori/go.uuid"
)

// Patient represents patients in the system
type Patient struct {
	ID            string   `json:"patient_id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	Prescriptions []string `json:"prescriptions"` // an array of prescription IDs for the patient
	Doctors       []string `json:"doctors"`
}

// NewPatient is the constructor for the Patient object
// this constructor is used for the regular user sign up flow
func NewPatient(name, email, phone string) *Patient {
	return &Patient{
		ID:            uuid.NewV4().String(), // FIXME: check unique
		Name:          name,
		Email:         email,
		Phone:         phone,
		Prescriptions: []string{},
		Doctors:       []string{},
	}
}

// NewPatientFromPrescription is the constructor for first time patients joining from a prescription
// this constructor is used when a user scans a QR code with the app
func NewPatientFromPrescription(name, phone, prescriptionID string) (*Patient, error) {
	p, err := prescriptions.GetPrescription(prescriptionID)
	if err != nil {
		return nil, prescriptions.ErrPrescriptionDoesNotExist
	}

	return &Patient{
		ID:            uuid.NewV4().String(), // FIXME: check unique
		Name:          name,
		Phone:         phone,
		Prescriptions: []string{prescriptionID},
		Doctors:       []string{p.Doctor},
	}, nil
}
