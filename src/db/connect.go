package db

import (
	"log"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var (
	// Session stores mongo session
	Session *mgo.Session

	// Mongo stores the mongodb connection string information
	Mongo *mgo.DialInfo
)

// Connect connects to mongodb
func Connect() {
	uri := os.Getenv("DB_URI")
	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal("Database Connectino: ", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	log.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}
