package doctors

import (
	"github.com/satori/go.uuid"
)

// Doctor represents a doctor
type Doctor struct {
	ID            string   `json:"doctor_id"`
	Name          string   `json:"name"`
	Office        string   `json:"office"`
	Specializaion string   `json:"specialization"`
	Patients      []string `json:"patients"`
}

// NewDoctor is the constructor for the Doctor object
func NewDoctor(name, office, specialization string) *Doctor {
	return &Doctor{
		ID:            uuid.NewV4().String(), // FIXME: check unique
		Name:          name,
		Office:        office,
		Specializaion: specialization,
		Patients:      []string{},
	}
}
