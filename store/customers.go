package store

import (
    "labix.org/v2/mgo/bson"
)

// TODO: define this as tags
var config = map[string][]string {
    "searchable": []string{
        "name",
        "email",
        "company",
        "address.city",
        "address.street",
    },
}

type Customer struct {
    Id          bson.ObjectId `bson:"_id"         json:"id"`
    Name        string        `bson:"name"        json:"name"`
    Email       string        `bson:"email"       json:"email"`
    Company     string        `bson:"company"     json:"company,omitempty"`
    Address     Address       `bson:"address"     json:"address"`
}

type Address struct {
    Street      string        `bson:"street"      json:"street,omitempty"`
    City        string        `bson:"city"        json:"city,omitempty"`
    Zip         uint16        `bson:"zip"         json:"zip,omitempty"`
}

func CustomersFind(query *Query) ([]Customer, error) {
    result := []Customer{}

    return result, db.C("customers").Find(query.Bson(config)).All(&result)
}

func CustomersFindOne(query *Query) (Customer, error) {
    result := Customer{}

    return result, db.C("customers").Find(query.Bson(config)).One(&result)
}
