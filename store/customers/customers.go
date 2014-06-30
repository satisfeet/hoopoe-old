package customers

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var (
	db *mgo.Database
)

type Customer struct {
	Id      bson.ObjectId   `bson:"_id"     json:"id"`
	Name    string          `bson:"name"    json:"name"`
	Email   string          `bson:"email"   json:"email"`
	Company string          `bson:"company" json:"company,omitempty"`
	Address CustomerAddress `bson:"address" json:"address"`
}

type CustomerAddress struct {
	Street string `bson:"street"  json:"street,omitempty"`
	City   string `bson:"city"    json:"city,omitempty"`
	Zip    uint16 `bson:"zip"     json:"zip,omitempty"`
}

func Setup(session *mgo.Session) {
	db = session.DB("")

	db.C("customers").EnsureIndex(mgo.Index{
		Key: []string{
			"email",
		},
		Unique: true,
	})
	db.C("customers").EnsureIndex(mgo.Index{
		Key: []string{
			"name",
			"company",
			"address.city",
			"address.street",
		},
	})
}

func FindAll(query *Query) ([]Customer, error) {
	r := []Customer{}

	return r, db.C("customers").Find(query).All(&r)
}

func FindOne(query *Query) (Customer, error) {
	r := Customer{}

	return r, db.C("customers").Find(query).One(&r)
}
