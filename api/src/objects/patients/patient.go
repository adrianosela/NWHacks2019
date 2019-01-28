package patients

import (
	"github.com/satori/go.uuid"
)

// Patient represents patients in the system
type Patient struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Age           int      `json:"age"`
	Gender        string   `json:"gender"`
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
	Age               int    `json:"age"`
	Gender            string `json:"gender,omitempty"`
	NewPrescriptionID string `json:"prescription_id,omitempty"` // empty if usual sign up flow
}

// NewPatient is the constructor for the Patient object
// this constructor is used for the regular user sign up flow
// whether new user is joining with a prescription or not
func NewPatient(config NewPatientConfig) *Patient {
	p := &Patient{
		ID:            uuid.Must(uuid.NewV4()).String(),
		Name:          config.Name,
		Email:         config.Email,
		Phone:         config.Phone,
		Age:           config.Age,
		Gender:        config.Gender,
		Prescriptions: []string{},
	}
	if config.NewPrescriptionID != "" {
		p.Prescriptions = []string{config.NewPrescriptionID}
	}
	return p
}
