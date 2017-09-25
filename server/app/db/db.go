package db

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Gateway struct {
	MAC      string `json:"mac" bson:"_id,omitempty"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
	Username string `json:"username"`
}

var db *mgo.Database

func Init(uri string, dbname string) error {
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db = session.DB(dbname)

	return nil
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
func Save(gw Gateway) (*mgo.ChangeInfo, error) {
	return collection().UpsertId(gw.MAC, gw)
}

// Remove deletes an gateway from the database
func Remove(id string) error {
	return collection().Remove(bson.M{"_id": id})
}
