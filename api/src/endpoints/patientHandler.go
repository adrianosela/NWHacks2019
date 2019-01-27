package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adrianosela/NWHacks2019/api/src/objects/doctors"
	"github.com/adrianosela/NWHacks2019/api/src/objects/patients"
	"github.com/adrianosela/NWHacks2019/api/src/objects/prescriptions"
	"github.com/adrianosela/NWHacks2019/api/src/store"
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
	// create patient object
	patient := patients.NewPatient(patReq)
	// if request contains prescription, tie doctor to patient
	if patReq.NewPrescriptionID != "" {
		var pres *prescriptions.Prescription
		pres, err = c.DB.GetPrescription(patReq.NewPrescriptionID)
		// fail closed if there was a problem with the provided prescription
		if err != nil {
			switch err {
			case store.ErrNotInStore:
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte(fmt.Sprintf("new patient request contained nonexistent prescription: %s", err.Error())))
				return
			default:
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(string(":(")))
				return
			}
		}

		//claim prescription for patiend and push change to db
		pres.Claimed = true
		pres.Patient = patient.ID
		if err = c.DB.UpdatePrescription(pres); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(string(":("))) //FIXME
			return
		}
		// tie doctor associated with prescription to the user
		patient.Doctors = []string{pres.Doctor}
		// add patient to doctor's db
		var dr *doctors.Doctor
		dr, err = c.DB.GetDoctor(pres.Doctor)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(string(":("))) //FIXME
			return
		}
		dr.Patients = append(dr.Patients, patient.ID)
		if err = c.DB.UpdateDoctor(dr); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(string(":("))) //FIXME
			return
		}

	}

	if err = c.DB.PutPatient(patient); err != nil {
		switch err {
		case store.ErrItemExists:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("UUID Collition occured"))
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
	p, err := c.DB.GetPatient(id)
	if err != nil {
		switch err {
		case store.ErrNotInStore:
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
