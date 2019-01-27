package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adrianosela/NWHacks2019/api/src/store"
	"github.com/gorilla/mux"
)

// APIConfig holds all necessary configurations for the service to run
type APIConfig struct {
	DB         store.DB
	DeployTime time.Time
}

// GetHandlers returns an HTTP MUX with the service's handler functions
func GetHandlers(c APIConfig) *mux.Router {
	router := mux.NewRouter()

	// doctors
	router.Methods(http.MethodPost).Path("/doctor").HandlerFunc(c.newDoctorHandler)
	router.Methods(http.MethodGet).Path("/doctor/{id}").HandlerFunc(c.getDoctorHandler)

	// patients
	router.Methods(http.MethodPost).Path("/patient").HandlerFunc(c.newPatientHandler)
	router.Methods(http.MethodGet).Path("/patient/{id}").HandlerFunc(c.getPatientHandler)
	router.Methods(http.MethodGet).Path("/patient_prescriptions/{id}").HandlerFunc(c.getPatientPrescriptionsHandler)
	router.Methods(http.MethodGet).Path("/patient_doctors/{id}").HandlerFunc(c.getPatientDoctorsHandler)

	// prescriptions
	router.Methods(http.MethodPost).Path("/prescription").HandlerFunc(c.newPrescriptionHandler)
	router.Methods(http.MethodGet).Path("/prescription/{id}").HandlerFunc(c.getPrescriptionHandler)
	router.Methods(http.MethodPost).Path("/claim").HandlerFunc(c.claimPrescriptionHandler)

	// test endpoint
	router.Methods(http.MethodGet).Path("/test").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("instance deployed at %s", c.DeployTime.String())))
	})

	return router
}
