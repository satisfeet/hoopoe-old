package store

import (
	"errors"
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/model/validation"
)

// Error which will be returned if the provided value does not have an id
// property.
var ErrInvalidType = errors.New("invalid type")

// Store builds onto Session and exposes high level CRUD operations.
//
// TODO: Actually it makes sense to merge session directly into store. However
// we still would need to solve the problem that the store keeps track of the
// collection name. Maybe we could use something like "Clone(name string)"?
type Store struct {
	Name    string
	Session *Session
}

// Returns a cloned mongo session.
//
// TODO: Remove this. This makes no sense at all.
func (s *Store) mongo() *mgo.Session {
	return s.Session.Mongo()
}

// Inserts a value into the database. If no id was defined it will auto-create
// one. Furthermore it will validate the given model before insertion. If any of
// these actions go wrong an error will be returned.
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

	if v, ok := v.(validation.Validatable); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return m.DB(Database).C(s.Name).Insert(v)
}

// Updates the given model matching the query with the given value. If query or
// model are not valid an error is returned.
//
// NOTE: Initially we did not want to use an extra Query parameter for this step
// however without it we would take heavy use of reflection and we would need to
// fetch the model before. With this way we save on database round trip.
func (s *Store) Update(q Query, v interface{}) error {
	m := s.mongo()
	defer m.Close()

	if !q.Valid() {
		return ErrInvalidQuery
	}

	if v, ok := v.(validation.Validatable); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return m.DB(Database).C(s.Name).Update(q, v)
}

// Removes the model matching the query from the database. If the query does not
// provide a valid id condition it will return an error or if it does not find a
// matching document.
func (s *Store) Remove(q Query) error {
	m := s.mongo()
	defer m.Close()

	if !q.Valid() {
		return ErrInvalidQuery
	}

	return m.DB(Database).C(s.Name).Remove(q)
}

// Maps all documents matching the given query to the given value.
func (s *Store) FindAll(q Query, v interface{}) error {
	m := s.mongo()
	defer m.Close()

	return m.DB(Database).C(s.Name).Find(q).All(v)
}

// Maps document with matching id to the given value.
func (s *Store) FindOne(q Query, v interface{}) error {
	m := s.mongo()
	defer m.Close()

	if !q.Valid() {
		return ErrInvalidQuery
	}

	return m.DB(Database).C(s.Name).Find(q).One(v)
}
