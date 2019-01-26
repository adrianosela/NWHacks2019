package endpoints

import "net/http"

func newDoctorHandler(w http.ResponseWriter, r *http.Request) {
	// TODO

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string("doctor account created")))
	return
}
