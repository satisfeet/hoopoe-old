package mongodb

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/satisfeet/hoopoe/store/common"
)

type Query bson.M

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

var DefaultStore = &Store{}

type Store struct {
	Session *mgo.Session
}

func (s *Store) Open(u string) error {
	c, err := mgo.Dial(u)
	if err != nil {
		return err
	}
	s.Session = c

	return nil
}

func (s *Store) Close() error {
	s.Session.Close()
	s.Session = nil

	return nil
}

func (s *Store) Find(q common.Query, m interface{}) error {
	c := s.Session.Clone()
	defer c.Close()

	return s.collection(m).With(c).Find(q).All(m)
}

func (s *Store) FindOne(q common.Query, m common.Model) error {
	c := s.Session.Clone()
	defer c.Close()

	return s.collection(m).With(c).Find(q).One(m)
}

func (s *Store) Insert(m common.Model) error {
	c := s.Session.Clone()
	defer c.Close()

	return s.collection(m).With(c).Insert(m)
}

func (s *Store) Update(m common.Model) error {
	c := s.Session.Clone()
	defer c.Close()

	return s.collection(m).With(c).UpdateId(common.GetId(m), m)
}

func (s *Store) Remove(m common.Model) error {
	c := s.Session.Clone()
	defer c.Close()

	return s.collection(m).With(c).RemoveId(common.GetId(m))
}

func (s *Store) collection(v interface{}) *mgo.Collection {
	return s.Session.DB("").C(common.GetName(v))
}
