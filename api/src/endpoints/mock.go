package endpoints

import (
	"net/http"
)

type mockRequest struct {
	// TODO
}

func mockHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(string("Hello world")))
	return
}
