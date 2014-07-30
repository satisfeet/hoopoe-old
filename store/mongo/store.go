package mongo

import (
	"errors"

	"gopkg.in/mgo.v2"
)

type Store struct {
	session *mgo.Session
}

var (
	ErrNotConnected   = errors.New("not connected")
	ErrStillConnected = errors.New("still connected")
)

func (s *Store) Dial(u string) error {
	if s.session != nil {
		return ErrStillConnected
	}

	sess, err := mgo.Dial(u)

	if err != nil {
		return err
	}

	s.session = sess

	return nil
}

func (s *Store) Close() error {
	if s.session == nil {
		return ErrNotConnected
	}

	s.session.Close()
	s.session = nil

	return nil
}

func (s *Store) clone() *mgo.Session {
	return s.session.Clone()
}

func (s *Store) filesystem(n string) *mgo.GridFS {
	return s.session.DB("").GridFS(n)
}

func (s *Store) collection(n string) *mgo.Collection {
	return s.session.DB("").C(n)
}

func (s *Store) Find(n string, q Query, v interface{}) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Find(q).All(v)
}

func (s *Store) FindOne(n string, q Query, v interface{}) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Find(q).One(v)
}

func (s *Store) Insert(n string, v interface{}) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Insert(v)
}

func (s *Store) Update(n string, q Query, v interface{}) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Update(q, v)
}

func (s *Store) Remove(n string, q Query) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Remove(q)
}

func (s *Store) RemoveAll(n string, q Query) error {
	c := s.clone()
	defer c.Close()

	_, err := s.collection(n).With(c).RemoveAll(q)

	return err
}
