package store

import "gopkg.in/mgo.v2"

// Default session used when no Session was passed to Store.
var session = &Session{}

// Opens the default session.
func Open(u string) error {
	return session.Open(u)
}

// Closes the default session.
func Close() error {
	return session.Close()
}

// Session is a container with database connections.
type Session struct {
	mongo *mgo.Session
}

// Opens all database connections.
func (s *Session) Open(u string) error {
	var err error
	s.mongo, err = mgo.Dial(u)

	return err
}

// Returns a copy of the mongodb connection and error if connection was not
// initialized.
func (s *Session) Mongo() (*mgo.Session, error) {
	if s.mongo == nil {
		return nil, ErrNotConnected
	}

	return s.mongo.Clone(), nil
}

// Closes all database connections.
func (s *Session) Close() error {
	if s.mongo == nil {
		return ErrNotConnected
	}

	s.mongo.Close()
	s.mongo = nil

	return nil
}
