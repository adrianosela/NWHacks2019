package endpoints

import "net/http"

func newPatientHandler(w http.ResponseWriter, r *http.Request) {
	// TODO

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string("patient account created")))
	return
}
