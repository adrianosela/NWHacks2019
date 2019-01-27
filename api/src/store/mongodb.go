package store

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/adrianosela/NWHacks2019/api/src/objects/doctors"
	"github.com/adrianosela/NWHacks2019/api/src/objects/patients"
	"github.com/adrianosela/NWHacks2019/api/src/objects/prescriptions"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoDB represents a Mongo DB instance
type MongoDB struct {
	Session            *mgo.Session
	DoctorsTable       *mgo.Collection
	PatientsTable      *mgo.Collection
	PrescriptionsTable *mgo.Collection
	MedicinesTable     *mgo.Collection
	Close              func()
}

// MongoDBConfig is used to configure a Mongo db instance
type MongoDBConfig struct {
	Host                   string
	Port                   int
	DBName                 string
	Timeout                time.Duration
	Username               string
	Password               string
	DoctorsTableName       string
	PatientsTableName      string
	PrescriptionsTableName string
	MedicinesTableName     string
}

// NewMongoDB returns a mock database which implements the DB interface
func NewMongoDB(c MongoDBConfig) (*MongoDB, error) {
	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{fmt.Sprintf("%s:%d", c.Host, c.Port)}, //"golang-couch.documents.azure.com:10255"
		Timeout:  c.Timeout,
		Database: c.DBName,
		Username: c.Username,
		Password: c.Password,
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}
	// create a session which maintains a pool of socket connections to our MongoDB.
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		return nil, fmt.Errorf("can't connect to mongo: %s", err)
	}

	// SetSafe changes the session safety mode.
	// If the safe parameter is nil, the session is put in unsafe mode, // and writes become fire-and-forget,
	// without error checking. The unsafe mode is faster since operations won't hold on waiting for a confirmation.
	// http://godoc.org/labix.org/v2/mgo#Session.SetMode.
	session.SetSafe(&mgo.Safe{})

	// return populated db
	return &MongoDB{
		Session:            session,
		DoctorsTable:       session.DB(c.DBName).C(c.DoctorsTableName),
		PatientsTable:      session.DB(c.DBName).C(c.PatientsTableName),
		PrescriptionsTable: session.DB(c.DBName).C(c.PrescriptionsTableName),
		MedicinesTable:     session.DB(c.DBName).C(c.MedicinesTableName),
		Close:              func() { session.Close() },
	}, nil
}

// PutDoctor stores a doctor in the db
func (db *MongoDB) PutDoctor(dr *doctors.Doctor) error {
	if _, err := db.GetDoctor(dr.ID); err == nil {
		return ErrItemExists
	}
	return db.DoctorsTable.Insert(dr)
}

// PutPatient stores a patient in the db
func (db *MongoDB) PutPatient(pt *patients.Patient) error {
	if _, err := db.GetPatient(pt.ID); err == nil {
		return ErrItemExists
	}
	return db.PatientsTable.Insert(pt)
}

// PutPrescription stores a prescription in the db
func (db *MongoDB) PutPrescription(rx *prescriptions.Prescription) error {
	if _, err := db.GetPrescription(rx.ID); err == nil {
		return ErrItemExists
	}
	return db.PrescriptionsTable.Insert(rx)
}

// UpdateDoctor updates a doctor in the db
func (db *MongoDB) UpdateDoctor(dr *doctors.Doctor) error {
	return db.DoctorsTable.Update(bson.M{"id": dr.ID}, dr)
}

// UpdatePatient updates a doctor in the db
func (db *MongoDB) UpdatePatient(pt *patients.Patient) error {
	return db.PatientsTable.Update(bson.M{"id": pt.ID}, pt)
}

// UpdatePrescription updates a prescription in the db
func (db *MongoDB) UpdatePrescription(pr *prescriptions.Prescription) error {
	return db.PrescriptionsTable.Update(bson.M{"id": pr.ID}, pr)
}

// GetDoctor gets a doctor from the db
func (db *MongoDB) GetDoctor(drID string) (*doctors.Doctor, error) {
	var dr *doctors.Doctor
	if err := db.DoctorsTable.Find(bson.M{"id": drID}).One(&dr); err != nil {
		return nil, fmt.Errorf("%s: %s", ErrNotInStore.Error(), err) //FIXME
	}
	return dr, nil
}

// GetPatient gets a patient from the db
func (db *MongoDB) GetPatient(ptID string) (*patients.Patient, error) {
	var pt *patients.Patient
	if err := db.PatientsTable.Find(bson.M{"id": ptID}).One(&pt); err != nil {
		return nil, fmt.Errorf("%s: %s", ErrNotInStore.Error(), err) //FIXME
	}
	return pt, nil
}

// GetPrescription gets a prescription from the db
func (db *MongoDB) GetPrescription(rxID string) (*prescriptions.Prescription, error) {
	var rx *prescriptions.Prescription
	if err := db.PrescriptionsTable.Find(bson.M{"id": rxID}).One(&rx); err != nil {
		return nil, fmt.Errorf("%s: %s", ErrNotInStore.Error(), err) //FIXME
	}
	return rx, nil
}
