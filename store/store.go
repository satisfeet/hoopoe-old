package store

import (
	"errors"
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrInvalidType = errors.New("Invalid type error.")
)

type Store struct {
	Name    string
	Session *Session
}

func (s *Store) mongo() *mgo.Session {
	if s.Session == nil {
		return DefaultSession.Mongo()
	}
	return s.Session.Mongo()
}

func (s *Store) Insert(v interface{}) error {
	m := s.mongo()
	defer m.Close()

	// set a bson object id
	if v := reflect.ValueOf(v).Elem().FieldByName("Id"); true {
		// check if id was not initialized so far
		if !v.Interface().(bson.ObjectId).Valid() {
			v.Set(reflect.ValueOf(bson.NewObjectId()))
		}
	} else {
		return ErrInvalidType
	}

	return m.DB(Database).C(s.Name).Insert(v)
}

func (s *Store) Update(q Query, v interface{}) error {
	m := s.mongo()
	defer m.Close()

	if !q.Valid() {
		return ErrInvalidQuery
	}

	return m.DB(Database).C(s.Name).Update(q, v)
}

func (s *Store) Remove(q Query) error {
	m := s.mongo()
	defer m.Close()

	if !q.Valid() {
		return ErrInvalidQuery
	}

	return m.DB(Database).C(s.Name).Remove(q)
}

func (s *Store) FindAll(q Query, v interface{}) error {
	m := s.mongo()
	defer m.Close()

	return m.DB(Database).C(s.Name).Find(q).All(v)
}

func (s *Store) FindOne(q Query, v interface{}) error {
	m := s.mongo()
	defer m.Close()

	if !q.Valid() {
		return ErrInvalidQuery
	}

	return m.DB(Database).C(s.Name).Find(q).One(v)
}
