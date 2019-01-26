package endpoints

import (
	"net/http"

	"github.com/gorilla/mux"
)

// GetHandlers returns an HTTP MUX with the service's handler functions
func GetHandlers() *mux.Router {
	router := mux.NewRouter()

	router.Methods(http.MethodGet).Path("/mock").HandlerFunc(mockHandler)

	return router
}
