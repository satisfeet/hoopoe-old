package store

import "labix.org/v2/mgo"

import "github.com/satisfeet/hoopoe/conf"

type Store struct {
	session *mgo.Session
}

func New() *Store {
	return &Store{}
}

func (s *Store) Open(c conf.Map) error {
	var err error

	s.session, err = mgo.Dial(c["mongo"])

	if err != nil {
		return err
	}

	return nil
}

func (s *Store) Manager(name string) *Manager {
	c := s.session.DB("").C(name)

	return &Manager{c}
}

func (s *Store) Close() {
	s.session.Close()

	return
}
