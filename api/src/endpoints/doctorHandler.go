package endpoints

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/adrianosela/NWHacks2019/api/src/objects/doctors"
	"github.com/adrianosela/NWHacks2019/api/src/store"
	"github.com/gorilla/mux"
)

func (c *APIConfig) newDoctorHandler(w http.ResponseWriter, r *http.Request) {
	// read req body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("could not read request body")))
		return
	}
	// unmarshal request payload
	var docReq doctors.NewDoctorConfig
	if err = json.Unmarshal(bodyBytes, &docReq); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("request is not of the correct format: " + err.Error())))
		return
	}
	// create doctor object
	d := doctors.NewDoctor(docReq)
	if err = c.DB.PutDoctor(d); err != nil {
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
	respBytes, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(string("could not marshall response")))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(respBytes)
	return
}

func (c *APIConfig) getDoctorHandler(w http.ResponseWriter, r *http.Request) {
	// get ID from URL params
	getParams := mux.Vars(r)
	id, ok := getParams["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(string("no id specified")))
		return
	}
	// get doctor from store
	d, err := c.DB.GetDoctor(id)
	if err != nil {
		switch err {
		case store.ErrNotInStore:
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("doctor %s not found", id)))
		default:
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(string(":(")))
			return
		}
	}
	// marshal doctor object
	responseBytes, err := json.Marshal(d)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(string("could not marshall response")))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseBytes)
	return
}
