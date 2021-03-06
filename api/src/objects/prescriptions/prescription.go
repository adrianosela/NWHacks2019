package prescriptions

import (
	"time"

	"github.com/satori/go.uuid"
)

// Prescription represents a prescription from a doctor to a patient
type Prescription struct {
	ID        string                 `json:"id"`
	AddedAt   int64                  `json:"added_at"`
	Medicines map[string]Indications `json:"medicines"` // map of medicine ID to indications
	Remaining map[string]int         `json:"remaining"` // map of medicine to quantity left
	Claimed   bool                   `json:"claimed"`
	Doctor    string                 `json:"doctor"`
	Patient   string                 `json:"patient,omitempty"`
}

// NewPrescriptionConfig represents all the necessary info to create a new prescription
type NewPrescriptionConfig struct {
	Medicines map[string]Indications `json:"medicines"`
	Amounts   map[string]int         `json:"amounts"`
	Patient   string                 `json:"patient,omitempty"`
	Doctor    string                 `json:"doctor"`
}

// ClaimPrescriptionConfig represents all the necessary info to claim an unclaimed prescription
type ClaimPrescriptionConfig struct {
	PatientID      string `json:"patient_id"`
	PrescriptionID string `json:"prescription_id"`
}

// NewPrescription adds a new prescription to the system
func NewPrescription(config NewPrescriptionConfig) *Prescription {
	p := &Prescription{
		ID:        uuid.Must(uuid.NewV4()).String(), // FIXME: check unique
		AddedAt:   time.Now().UnixNano(),
		Medicines: config.Medicines,
		Remaining: config.Amounts,
		Doctor:    config.Doctor,
	}

	if config.Patient != "" {
		p.Claimed = true
		p.Patient = config.Patient
	}

	return p
}
