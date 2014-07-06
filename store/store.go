package store

import (
	"labix.org/v2/mgo"

	"github.com/satisfeet/hoopoe/conf"
)

type Store struct {
	mongo *mgo.Session
}

func New() *Store {
	return &Store{}
}

func (s *Store) Open(c conf.Map) error {
	var err error

	s.mongo, err = mgo.Dial(c["mongo"])

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Mongo() *mgo.Session {
	return s.mongo
}

func (s *Store) Close() {
	s.mongo.Close()
}
