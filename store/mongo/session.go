package mongo

import "gopkg.in/mgo.v2"

// The Session type represents a connection to a mongodb server. It is used to
// have an unified interface for all database sessions in common.
type Session struct {
	session  *mgo.Session
	database *mgo.Database
}

func (s *Session) Dial(url string) error {
	session, err := mgo.Dial(url)

	if err != nil {
		return err
	}

	s.session = session
	s.database = session.DB("")

	return nil
}

func (s *Session) Close() error {
	s.session.Close()

	s.session = nil
	s.database = nil

	return nil
}
