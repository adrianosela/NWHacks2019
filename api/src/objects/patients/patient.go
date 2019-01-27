package patients

import (
	"github.com/adrianosela/NWHacks2019/api/src/objects/prescriptions"
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

// NewPatientConfig represents all the necessary info to create a new patient
type NewPatientConfig struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	NewPrescriptionID string `json:"prescription_id,omitempty"` // empty if usual sign up flow
}

// NewPatient is the constructor for the Patient object
// this constructor is used for the regular user sign up flow
// whether new user is joining with a prescription or not
func NewPatient(config NewPatientConfig) (*Patient, error) {
	np := &Patient{
		ID:    uuid.NewV4().String(), // FIXME: check unique
		Name:  config.Name,
		Email: config.Email,
		Phone: config.Phone,
	}
	// if config contains a prescription, link it and the doctor to the new
	// patient's account
	if config.NewPrescriptionID != "" {
		pres, err := prescriptions.GetPrescription(config.NewPrescriptionID)
		if err != nil {
			return nil, err
		}
		np.Prescriptions = []string{config.NewPrescriptionID}
		np.Doctors = []string{pres.Doctor}
	}
	// FIXME: STORE PATIENT IN DB
	return np, nil
}

// GetPatient returns a patient given the patient ID
func GetPatient(id string) (*Patient, error) {
	// FIXME: GET PATIENT FROM DB
	return &Patient{
		ID: id, // FIXME: check unique
	}, nil
}
