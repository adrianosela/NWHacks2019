package endpoints

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/adrianosela/NWHacks2019/api/src/objects/prescriptions"
	"github.com/gorilla/mux"
)

func newPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	// read req body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("could not read request body")))
		return
	}
	// unmarshal request payload
	var rxReq prescriptions.NewPrescriptionConfig
	if err = json.Unmarshal(bodyBytes, &rxReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("request is not of the correct format: " + err.Error())))
		return
	}
	// create new prescription object
	p := prescriptions.NewPrescription(rxReq)
	// marshal response payload
	respBytes, err := json.Marshal(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(string("could not marshall response")))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}

func getPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	// get prescription id from URL
	getParams := mux.Vars(r)
	id, ok := getParams["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("no id specified")))
		return
	}
	// get prescription from store
	p, err := prescriptions.GetPrescription(id)
	if err != nil {
		switch err {
		case prescriptions.ErrPrescriptionDoesNotExist:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(string(":(")))
			return
		}
	}
	// marshal response
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
