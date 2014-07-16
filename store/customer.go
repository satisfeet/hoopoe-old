package store

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	CustomerUnique = []string{
		"email",
	}

	CustomerIndices = []string{
		"name",
		"company",
		"address.street",
		"address.city",
	}
)

type Customer struct {
	Id      bson.ObjectId   `json:"id"     bson:"_id"`
	Name    string          `json:"name"`
	Email   string          `json:"email"`
	Company string          `json:"company,omitempty"`
	Address CustomerAddress `json:"address"`
}

type CustomerAddress struct {
	Zip    int    `json:"zip,omitempty"`
	City   string `json:"city,omitempty"`
	Street string `json:"street,omitempty"`
}

func (s *Store) IndexCustomer() {
	c := s.mongo.DB("").C("customers")

	c.EnsureIndex(mgo.Index{Key: CustomerIndices})
	c.EnsureIndex(mgo.Index{Key: CustomerUnique, Unique: true})
}

func (s *Store) InsertCustomer(c *Customer) error {
	m := s.mongo.Clone()
	defer m.Close()

	if !c.Id.Valid() {
		c.Id = bson.NewObjectId()
	}

	return m.DB("").C("customers").Insert(c)
}

func (s *Store) UpdateCustomer(c *Customer) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB("").C("customers").UpdateId(c.Id, c)
}

func (s *Store) RemoveCustomer(q Query) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB("").C("customers").Remove(q)
}

func (s *Store) FindAllCustomer(q Query, c *[]Customer) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB("").C("customers").Find(q).All(c)
}

func (s *Store) FindOneCustomer(q Query, c *Customer) error {
	m := s.mongo.Clone()
	defer m.Close()

	return m.DB("").C("customers").Find(q).One(c)
}
