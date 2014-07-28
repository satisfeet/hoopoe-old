package mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store/common"
)

// Query abstracts conditions for mongocommon.
type Query bson.M

// Adds an equals id condition to query.
func (q Query) Id(v interface{}) error {
	switch t := v.(type) {
	case string:
		if bson.IsObjectIdHex(t) {
			q["_id"] = bson.ObjectIdHex(t)

			return nil
		}
		return common.ErrBadQueryId
	case bson.ObjectId:
		if t.Valid() {
			q["_id"] = t

			return nil
		}
		return common.ErrBadQueryId
	}

	return common.ErrBadQueryValue
}

// Store abstracts a store backed up by mongocommon.
type Store struct {
	session *mgo.Session
}

// Establishes a connection to database.
func (s *Store) Open(u string) error {
	if s.session != nil {
		return common.ErrStillConnected
	}

	sess, err := mgo.Dial(u)
	if err != nil {
		return err
	}

	s.session = sess

	return nil
}

// Closes established connection.
func (s *Store) Close() error {
	if s.session == nil {
		return common.ErrNotConnected
	}
	s.session.Close()
	s.session = nil

	return nil
}

// Allocates new mongo session.
func (s *Store) clone() *mgo.Session {
	return s.session.Clone()
}

// Returns collection with name.
func (s *Store) collection(n string) *mgo.Collection {
	return s.session.DB("").C(n)
}

// Drops collection of documents.
func (s *Store) Drop(n string) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).DropCollection()
}

// Inserts document with value into collection.
func (s *Store) Insert(n string, v interface{}) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Insert(v)
}

// Updates document matching query with value.
func (s *Store) Update(n string, q Query, v interface{}) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Update(q, v)
}

// Removes document matching query from collection.
func (s *Store) Remove(n string, q Query) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Remove(q)
}

// Maps documents matching query onto interface.
func (s *Store) FindAll(n string, q Query, v interface{}) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Find(q).All(v)
}

// Maps document matching query onto interface.
func (s *Store) FindOne(n string, q Query, v interface{}) error {
	c := s.clone()
	defer c.Close()

	return s.collection(n).With(c).Find(q).One(v)
}
