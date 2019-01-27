package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adrianosela/NWHacks2019/api/src/objects/prescriptions"
	"github.com/adrianosela/NWHacks2019/api/src/store"
	"github.com/gorilla/mux"
)

func (c *APIConfig) newPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
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
	if err = c.DB.PutPrescription(p); err != nil {
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

func (c *APIConfig) getPrescriptionHandler(w http.ResponseWriter, r *http.Request) {
	// get prescription id from URL
	getParams := mux.Vars(r)
	id, ok := getParams["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("no id specified")))
		return
	}
	// get prescription from store
	p, err := c.DB.GetPrescription(id)
	if err != nil {
		switch err {
		case store.ErrNotInStore:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("prescription %s not found", id)))
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
