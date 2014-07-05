package store

import "labix.org/v2/mgo/bson"

type Customer struct {
	Oid     bson.ObjectId   `bson:"_id"     json:"id"`
	Name    string          `bson:"name"    json:"name"              store:"index"`
	Email   string          `bson:"email"   json:"email"             store:"unique"`
	Company string          `bson:"company" json:"company,omitempty" store:"index"`
	Address CustomerAddress `bson:"address" json:"address"`
}

type CustomerAddress struct {
	Street string `bson:"street"  json:"street,omitempty" store:"index"`
	City   string `bson:"city"    json:"city,omitempty"   store:"index"`
	Zip    uint16 `bson:"zip"     json:"zip,omitempty"`
}

func (c *Customer) Id() bson.ObjectId {
	return c.Oid
}

func (c *Customer) NewId() {
	c.Oid = bson.NewObjectId()
}
