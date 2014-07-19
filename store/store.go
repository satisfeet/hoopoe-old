package store

import (
	"errors"
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Store struct {
	// The mongo collection to operate on.
	mongo *mgo.Collection
}

var (
	InvalidTypeError  = errors.New("Invalid type error.")
	InvalidQueryError = errors.New("Invalid query error.")
)

func NewStore(n string) *Store {
	m := mongo.DB(Database).C(n)

	return &Store{
		mongo: m,
	}
}

func (s *Store) Insert(v interface{}) error {
	m := mongo.Clone()
	defer m.Close()

	// set a bson object id
	if v := reflect.ValueOf(v).Elem().FieldByName("Id"); true {
		// check if id was not initialized so far
		if !v.Interface().(bson.ObjectId).Valid() {
			v.Set(reflect.ValueOf(bson.NewObjectId()))
		}
	} else {
		return InvalidTypeError
	}

	return s.mongo.With(m).Insert(v)
}

func (s *Store) Update(q Query, v interface{}) error {
	m := mongo.Clone()
	defer m.Close()

	if !q.Valid() {
		return InvalidQueryError
	}

	return s.mongo.With(m).Update(q, v)
}

func (s *Store) Remove(q Query) error {
	m := mongo.Clone()
	defer m.Close()

	if !q.Valid() {
		return InvalidQueryError
	}

	return s.mongo.With(m).Remove(q)
}

func (s *Store) FindAll(q Query, v interface{}) error {
	m := mongo.Clone()
	defer m.Close()

	return s.mongo.With(m).Find(q).All(v)
}

func (s *Store) FindOne(q Query, v interface{}) error {
	m := mongo.Clone()
	defer m.Close()

	if !q.Valid() {
		return InvalidQueryError
	}

	return s.mongo.With(m).Find(q).One(v)
}
