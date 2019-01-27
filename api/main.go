package main

import (
	"log"
	"net/http"

	"github.com/adrianosela/NWHacks2019/api/src/endpoints"
	"github.com/adrianosela/NWHacks2019/api/src/store"
)

func main() {
	h := endpoints.GetHandlers(endpoints.APIConfig{
		DB: store.NewMockDB(),
	})
	err := http.ListenAndServe(":80", h)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
