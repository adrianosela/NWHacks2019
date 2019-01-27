package main

import (
	"log"
	"net/http"
	"time"

	"github.com/adrianosela/NWHacks2019/api/src/endpoints"
	"github.com/adrianosela/NWHacks2019/api/src/store"
)

func main() {
	h := endpoints.GetHandlers(endpoints.APIConfig{
		DB:         store.NewMockDB(),
		DeployTime: time.Now(),
	})
	err := http.ListenAndServe(":80", h)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
