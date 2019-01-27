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

// NewDoctorConfig represents all the necessary info to create a new doctor
type NewDoctorConfig struct {
	Name           string `json:"name"`
	Office         string `json:"office"`
	Specialization string `json:"specialization"`
}

// NewDoctor is the constructor for the Doctor object
func NewDoctor(config NewDoctorConfig) (*Doctor, error) {
	return &Doctor{
		ID:            uuid.NewV4().String(), // FIXME: check unique
		Name:          config.Name,
		Office:        config.Office,
		Specializaion: config.Specialization,
		Patients:      []string{},
	}, nil
	// FIXME: STORE IN DB AND CHECK UNIQUE
}

// GetDoctor returns a doctor given the doctor ID
func GetDoctor(id string) (*Doctor, error) {
	// FIXME: GET DOCTOR FROM DB
	return &Doctor{
		ID: id,
	}, nil
}
