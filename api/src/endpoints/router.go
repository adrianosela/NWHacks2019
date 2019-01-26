package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GetHandlers returns an HTTP MUX with the service's handler functions
func GetHandlers() *mux.Router {
	router := mux.NewRouter()

	// doctors
	router.Methods(http.MethodPost).Path("/doctor").HandlerFunc(newDoctorHandler)

	// patients
	router.Methods(http.MethodPost).Path("/patient").HandlerFunc(newPatientHandler)

	// prescriptions
	router.Methods(http.MethodPost).Path("/prescription").HandlerFunc(newPrescriptionHandler)

	return router
}
