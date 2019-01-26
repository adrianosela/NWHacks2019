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
	router.Methods(http.MethodGet).Path("/doctor/{id}").HandlerFunc(getDoctorHandler)

	// patients
	router.Methods(http.MethodPost).Path("/patient").HandlerFunc(newPatientHandler)
	router.Methods(http.MethodGet).Path("/patient/{id}").HandlerFunc(getPatientHandler)

	// prescriptions
	router.Methods(http.MethodPost).Path("/prescription").HandlerFunc(newPrescriptionHandler)
	router.Methods(http.MethodGet).Path("/prescription/{id}").HandlerFunc(getPrescriptionHandler)

	return router
}
