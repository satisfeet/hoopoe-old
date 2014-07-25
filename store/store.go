package store

import (
	"errors"
	"reflect"

	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model/validation"
)

var (
	ErrNotConnected = errors.New("not connected")
	ErrInvalidQuery = errors.New("invalid query")
	ErrInvalidType  = errors.New("invalid type")
)

// Store provides high-level CRUD interface to databases.
type Store struct {
	Name    string
	Session *Session
}

// Returns provided Session or global as fallback.
func (s *Store) session() *Session {
	if s.Session == nil {
		return session
	}
	return s.Session
}

func (s *Store) Insert(v interface{}) error {
	m, err := s.session().Mongo()
	if err != nil {
		return err
	}
	defer m.Close()

	if v := reflect.ValueOf(v).Elem().FieldByName("Id"); true {
		if !v.Interface().(bson.ObjectId).Valid() {
			v.Set(reflect.ValueOf(bson.NewObjectId()))
		}
	} else {
		return ErrInvalidType
	}

	if v, ok := v.(validation.Validatable); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return m.DB("").C(s.Name).Insert(v)
}

func (s *Store) Update(q Query, v interface{}) error {
	m, err := s.session().Mongo()
	if err != nil {
		return err
	}
	defer m.Close()

	if !q.Valid() {
		return ErrInvalidQuery
	}

	if v, ok := v.(validation.Validatable); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return m.DB("").C(s.Name).Update(q, v)
}

func (s *Store) Remove(q Query) error {
	m, err := s.session().Mongo()
	if err != nil {
		return err
	}
	defer m.Close()

	if !q.Valid() {
		return ErrInvalidQuery
	}

	return m.DB("").C(s.Name).Remove(q)
}

func (s *Store) FindAll(q Query, v interface{}) error {
	m, err := s.session().Mongo()
	if err != nil {
		return err
	}
	defer m.Close()

	return m.DB("").C(s.Name).Find(q).All(v)
}

func (s *Store) FindOne(q Query, v interface{}) error {
	m, err := s.session().Mongo()
	if err != nil {
		return err
	}
	defer m.Close()

	if !q.Valid() {
		return ErrInvalidQuery
	}

	return m.DB("").C(s.Name).Find(q).One(v)
}
