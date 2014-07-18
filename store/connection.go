package store

import "gopkg.in/mgo.v2"

var (
	// The global mongodb session.
	mongo *mgo.Session = nil
)

const (
	// Uses default database defined by mongo url.
	Database = ""
)

// Opens all database connections.
func Open(s string) error {
	var err error
	mongo, err = mgo.Dial(s)

	if err != nil {
		return err
	}
	return nil
}

// Closes all database connections.
func Close() {
	mongo.Close()
	mongo = nil
}
