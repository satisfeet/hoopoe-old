package customers

import (
	"labix.org/v2/mgo/bson"
)

type Model struct {
	Id          bson.ObjectId	`bson:"_id"         json:"id"`
	Name        string			`bson:"name"        json:"name"`
	Email       string			`bson:"email"       json:"email"`
	Company     string			`bson:"company"     json:"company,omitempty"`
	Address     ModelAddress	`bson:"address"     json:"address"`
}

type ModelAddress struct {
	Street      string        `bson:"street"      json:"street,omitempty"`
	City        string        `bson:"city"        json:"city,omitempty"`
	Zip         uint16        `bson:"zip"         json:"zip,omitempty"`
}
