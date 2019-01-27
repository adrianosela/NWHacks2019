package endpoints

import (
	"net/http"

	"github.com/adrianosela/NWHacks2019/api/src/store"
	"github.com/gorilla/mux"
)

// APIConfig holds all necessary configurations for the service to run
type APIConfig struct {
	DB store.DB
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

	// prescriptions
	router.Methods(http.MethodPost).Path("/prescription").HandlerFunc(c.newPrescriptionHandler)
	router.Methods(http.MethodGet).Path("/prescription/{id}").HandlerFunc(c.getPrescriptionHandler)

	return router
}
