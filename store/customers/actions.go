package customers

import (
    "labix.org/v2/mgo"
)

var (
	db *mgo.Database
)

func Open(session *mgo.Session) {
	db = session.DB("")
}

func FindAll(query *Query) ([]Model, error) {
    r := []Model{}

    return r, db.C("customers").Find(query).All(&r)
}

func  FindOne(query *Query) (Model, error) {
    r := Model{}

    return r, db.C("customers").Find(query).One(&r)
}
