package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/adrianosela/NWHacks2019/api/src/endpoints"
	"github.com/adrianosela/NWHacks2019/api/src/store"
)

func main() {
	db, err := store.NewMongoDB(store.MongoDBConfig{
		Host:                   "nwhacks2019.documents.azure.com",
		Port:                   10255,
		Username:               "nwhacks2019",
		Password:               os.Getenv("MONGODB_PASSWORD"),
		Timeout:                time.Second * 60,
		DBName:                 "nwhacksdb",
		DoctorsTableName:       "Doctors",
		PatientsTableName:      "Patients",
		PrescriptionsTableName: "Prescriptions",
	})
	defer db.Close()
	if err != nil {
		log.Fatal(err) // can't start without a db
	}

	h := endpoints.GetHandlers(endpoints.APIConfig{
		DB:         db,
		DeployTime: time.Now(),
	})

	if err = http.ListenAndServe(":80", h); err != nil {
		log.Fatal("ListenAndServe Error: ", err)
	}
}
