package endpoints

import "net/http"

func newPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	// TODO

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string("prescription created")))
	return
}
