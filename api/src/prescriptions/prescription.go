package prescriptions

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Prescription represents a prescription from a doctor to a patient
type Prescription struct {
	ID        string                 `json:"rx_id"`
	AddedAt   int64                  `json:"added_at"`
	Medicines map[string]Indications `json:"meds"` // map of medicine ID to indications
	Remaining map[string]int         `json:"remaining"`
	Claimed   bool                   `json:"claimed"`
	Patient   string                 `json:"patient"`
}

// Indications represents the frequency and way of ingesting the meds
type Indications struct {
	DaysPerWeek int   `json:"days_per_week"`
	TimesPerDay int   `json:"times_per_day"`
	Time        []int `json:"times_of_day"` // HOUR OF DAY FMT E.G. 1100, 2359, etc
}

// NewPrescription adds a new prescription to the system
func NewPrescription(medicines map[string]Indications, amounts map[string]int, patient string) *Prescription {
	p := &Prescription{
		ID:        uuid.NewV4().String(), // FIXME: check unique
		AddedAt:   time.Now().UnixNano(),
		Medicines: medicines,
		Remaining: amounts,
	}

	if patient != "" {
		p.Claimed = true
		p.Patient = patient
	}

	return p
}

// GetPrescription returns prescription details given a prescription id
func GetPrescription(id string) *Prescription {
	// TODO: hook DB here
	return &Prescription{
		ID: id,
	}
}
