package main

import (
	"log"
	"net/http"

	"github.com/adrianosela/NWHacks2019/api/endpoints"
)

func main() {
	h := endpoints.GetHandlers()
	err := http.ListenAndServe(":80", h)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
