package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newDoctorHandler(w http.ResponseWriter, r *http.Request) {
	// TODO

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string("doctor account created")))
	return
}

func getDoctorHandler(w http.ResponseWriter, r *http.Request) {
	getParams := mux.Vars(r)

	id, ok := getParams["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("no id specified")))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("here's doctor %s", id)))
	return
}
