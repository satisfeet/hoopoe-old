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

func IndexCustomer(s *Store) {
	db := s.Mongo()
	defer db.Session.Close()

	c := db.C("customers")

	c.EnsureIndex(mgo.Index{Key: CustomerIndices})
	c.EnsureIndex(mgo.Index{Key: CustomerUnique, Unique: true})
}

func InsertCustomer(s *Store, c *Customer) error {
	db := s.Mongo()
	defer db.Session.Close()

	if !c.Id.Valid() {
		c.Id = bson.NewObjectId()
	}

	return db.C("customers").Insert(c)
}

func UpdateCustomer(s *Store, c *Customer) error {
	db := s.Mongo()
	defer db.Session.Close()

	return db.C("customers").UpdateId(c.Id, c)
}

func RemoveCustomer(s *Store, q Query) error {
	db := s.Mongo()
	defer db.Session.Close()

	return db.C("customers").Remove(q)
}

func FindAllCustomer(s *Store, q Query, c *[]Customer) error {
	db := s.Mongo()
	defer db.Session.Close()

	return db.C("customers").Find(q).All(c)
}

func FindOneCustomer(s *Store, q Query, c *Customer) error {
	db := s.Mongo()
	defer db.Session.Close()

	return db.C("customers").Find(q).One(c)
}
