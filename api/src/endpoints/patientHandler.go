package endpoints

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func newPatientHandler(w http.ResponseWriter, r *http.Request) {
	// TODO

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string("patient account created")))
	return
}

func getPatientHandler(w http.ResponseWriter, r *http.Request) {
	getParams := mux.Vars(r)
	id, ok := getParams["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("no id specified")))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("here's patient %s", id)))
	return
}
