package store

import (
	"errors"

	"gopkg.in/mgo.v2"
)

var (
	// As we define the database to use in the mongo url we use this empty
	// placeholder. To avoid code empty strings everywhere we use this
	// variable as placeholder.
	Database = ""

	// Error returned when no database sessions where established.
	// This can be simply solved by calling Open on a session before use.
	ErrNotConnected = errors.New("not connected")
)

// Session is a database connection container. It is used as dependency
// for higher level structs to provide an unified access to multiple databases.
//
// NOTE: At the moment this is mongodb only though it should be easy to add
// more databases with time.
type Session struct {
	mongo *mgo.Session
}

// Opens all databases with the given parameter. Returns an error if something
// went wrong.
//
// NOTE: You NEED to call this before any further database actions can be
// performed else you will get ErrNotConnected.
func (s *Session) Open(u string) error {
	var err error
	s.mongo, err = mgo.Dial(u)
	return err
}

// Returns a cloned connection to the mongodb service. This will panic if no
// connection was opened before.
//
// TODO: Do not expose the raw mongo session to the public. Just not sure how
// to hide this else?
func (s *Session) Mongo() *mgo.Session {
	if s.mongo == nil {
		panic(ErrNotConnected)
	}

	return s.mongo.Clone()
}

// Closes all database connections or panics if not connected.
func (s *Session) Close() {
	if s.mongo == nil {
		panic(ErrNotConnected)
	}

	s.mongo.Close()
	s.mongo = nil
}
