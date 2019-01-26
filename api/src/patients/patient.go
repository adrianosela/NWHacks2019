package patients

import uuid "github.com/satori/go.uuid"

// Patient represents patients in the system
type Patient struct {
	ID            string   `json:"patient_id"`
	Name          string   `json:"name"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	Prescriptions []string `json:"prescriptions"` // an array of prescription IDs for the patient
}

// NewPatient is the constructor for the Patient object
func NewPatient(name, email, phone string) *Patient {
	return &Patient{
		ID:            uuid.NewV4().String(), // FIXME: check unique
		Name:          name,
		Email:         email,
		Phone:         phone,
		Prescriptions: []string{},
	}
}

// NewPatientFromPrescription is the constructor for first time patients joining from a prescription
func NewPatientFromPrescription(name, phone, prescriptionID string) *Patient {
	return &Patient{
		ID:            uuid.NewV4().String(), // FIXME: check unique
		Name:          name,
		Phone:         phone,
		Prescriptions: []string{prescriptionID},
	}
}
