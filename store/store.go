package store

import "gopkg.in/mgo.v2"

type Store struct {
	// The mongo collection to operate on.
	mongo *mgo.Collection
}

func NewStore(n string) *Store {
	m := mongo.DB(Database).C(n)

	return &Store{
		mongo: m,
	}
}

func (s *Store) Insert(v interface{}) error {
	m := mongo.Clone()
	defer m.Close()

	return s.mongo.With(m).Insert(v)
}

func (s *Store) Update(q Query, v interface{}) error {
	m := mongo.Clone()
	defer m.Close()

	return s.mongo.With(m).Update(q, v)
}

func (s *Store) Remove(q Query) error {
	m := mongo.Clone()
	defer m.Close()

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

	return s.mongo.With(m).Find(q).One(v)
}
