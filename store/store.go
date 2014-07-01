package store

import "labix.org/v2/mgo"

var (
	db *mgo.Database
)

// Alias to condition map.
type Query map[string]string

// Open all database connections.
func Open(c map[string]string) error {
	s, err := mgo.Dial(c["mongo"])

	if err != nil {
		return err
	}

	db = s.DB("")

	CustomersIndex()

	return nil
}

// Close all database connections.
func Close() {
	db.Session.Close()
}
