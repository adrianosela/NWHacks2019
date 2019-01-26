package prescriptions

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Prescription represents a prescription from a doctor to a patient
type Prescription struct {
	ID        string                 `json:"rx_id"`
	AddedAt   int64                  `json:"added_at"`
	Medicines map[string]Indications `json:"medicines"` // map of medicine ID to indications
	Remaining map[string]int         `json:"remaining"` // map of medicine to quantity left
	Claimed   bool                   `json:"claimed"`
	Doctor    string                 `json:"doctor"`
	Patient   string                 `json:"patient,omitempty"`
}

// Indications represents the frequency and way of ingesting the meds
type Indications struct {
	DaysPerWeek int   `json:"days_per_week,omitempty"`
	TimesPerDay int   `json:"times_per_day,omitempty"`
	Time        []int `json:"times_of_day,omitempty"` // HOUR OF DAY FMT E.G. 1100, 2359, etc
}

// NewPrescriptionConfig represents all the necessary info to create a new prescription
type NewPrescriptionConfig struct {
	Medicines map[string]Indications `json:"medicines"`
	Amounts   map[string]int         `json:"amounts"`
	Patient   string                 `json:"patient,omitempty"`
	Doctor    string                 `json:"doctor"`
}

// NewPrescription adds a new prescription to the system
func NewPrescription(config NewPrescriptionConfig) *Prescription {
	p := &Prescription{
		ID:        uuid.NewV4().String(), // FIXME: check unique
		AddedAt:   time.Now().UnixNano(),
		Medicines: config.Medicines,
		Remaining: config.Amounts,
		Doctor:    config.Doctor,
	}

	if config.Patient != "" {
		p.Claimed = true
		p.Patient = config.Patient
	}

	// FIXME: STORE PRESCRIPTION IN DB

	return p
}

// GetPrescription returns prescription details given a prescription id
func GetPrescription(id string) *Prescription {
	// TODO: hook DB here
	return &Prescription{
		AddedAt: time.Now().UnixNano(),
		ID:      id,
	}
}
