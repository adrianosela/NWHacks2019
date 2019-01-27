package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adrianosela/NWHacks2019/api/src/objects/patients"
	"github.com/adrianosela/NWHacks2019/api/src/objects/prescriptions"
	"github.com/gorilla/mux"
)

func (c *APIConfig) newPatientHandler(w http.ResponseWriter, r *http.Request) {
	// read req body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("could not read request body")))
		return
	}
	// unmarshal request payload
	var patReq patients.NewPatientConfig
	if err = json.Unmarshal(bodyBytes, &patReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("request is not of the correct format: " + err.Error())))
		return
	}
	// create patient in store
	patient, err := patients.NewPatient(patReq)
	if err != nil {
		switch err {
		case prescriptions.ErrPrescriptionDoesNotExist:
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(string("request included prescription but it does not exist")))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(string(":(")))
			return
		}
	}
	// marshal response payload
	respBytes, err := json.Marshal(patient)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(string("could not marshall response")))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}

func (c *APIConfig) getPatientHandler(w http.ResponseWriter, r *http.Request) {
	// get ID from URL params
	getParams := mux.Vars(r)
	id, ok := getParams["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("no id specified")))
		return
	}
	// get patient from store
	p, err := patients.GetPatient(id)
	if err != nil {
		switch err {
		case patients.ErrPatientNotFound:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("patient %s not found", id)))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(string(":(")))
			return
		}
	}
	// marshal patient object
	responseBytes, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(string("could not marshall response")))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
	return
}
