package store

import (
    "labix.org/v2/mgo/bson"
)

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

func CustomersFind() ([]Customer, error) {
    result := []Customer{}

    err := db.C("customers").Find(bson.M{}).All(&result)

    return result, err
}

func CustomersFindOne(param string) (Customer, error) {
    result := Customer{}

    query := bson.M{
        "_id": bson.ObjectIdHex(param),
    }

    err := db.C("customers").Find(query).One(&result)

    return result, err
}
