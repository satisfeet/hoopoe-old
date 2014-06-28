package customers

import (
    "labix.org/v2/mgo"
)

var (
	db *mgo.Database
)

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

func  FindOne(query *Query) (Customer, error) {
    r := Customer{}

    return r, db.C("customers").Find(query).One(&r)
}
