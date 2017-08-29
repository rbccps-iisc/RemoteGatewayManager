package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Gateway struct {
	MAC string `json:"mac" bson:"_id,omitempty"`
	IP  string `json:"ip"`
}

var db *mgo.Database

func init() {
	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB("api_db")
}

func collection() *mgo.Collection {
	return db.C("gws")
}

// GetAll returns all gateways from the database.
func GetAll() ([]Gateway, error) {
	res := []Gateway{}

	if err := collection().Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil
}

// GetOne returns a single gateway from the database.
func GetOne(id string) (*Gateway, error) {
	res := Gateway{}

	if err := collection().Find(bson.M{"_id": id}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

// Save inserts an gateway to the database.
func Save(gw Gateway) error {
	return collection().Insert(gw)
}

// Remove deletes an gateway from the database
func Remove(id string) error {
	return collection().Remove(bson.M{"_id": id})
}
