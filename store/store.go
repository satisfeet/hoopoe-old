package store

import "labix.org/v2/mgo"

type Store struct {
	mongo *mgo.Session
}

func New() *Store {
	return &Store{}
}

func (s *Store) Open(url string) error {
	var err error

	s.mongo, err = mgo.Dial(url)

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Mongo() *mgo.Session {
	if s.mongo == nil {
		panic("You need open store before!")
	}

	return s.mongo
}

func (s *Store) Close() {
	s.mongo.Close()

	s.mongo = nil
}
