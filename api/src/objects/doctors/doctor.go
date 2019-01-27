package doctors

import (
	"github.com/satori/go.uuid"
)

// Doctor represents a doctor
type Doctor struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Office        string   `json:"office"`
	OfficePhone   string   `json:"office_phone,omitempty"`
	PhotoURL      string   `json:"photo_url,omitempty"`
	Specializaion string   `json:"specialization"`
	Patients      []string `json:"patients"`
}

// NewDoctorConfig represents all the necessary info to create a new doctor
type NewDoctorConfig struct {
	Name           string `json:"name"`
	Office         string `json:"office"`
	OfficePhone    string `json:"office_phone,omitempty"`
	PhotoURL       string `json:"photo_url,omitempty"`
	Specialization string `json:"specialization"`
}

// NewDoctor is the constructor for the Doctor object
func NewDoctor(config NewDoctorConfig) *Doctor {
	return &Doctor{
		ID:            uuid.NewV4().String(),
		Name:          config.Name,
		Office:        config.Office,
		OfficePhone:   config.OfficePhone,
		Specializaion: config.Specialization,
		PhotoURL:      config.PhotoURL,
		Patients:      []string{},
	}
}
