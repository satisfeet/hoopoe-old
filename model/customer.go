package model

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

type Customer struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `index:"unique" validate:"required,min=5"`
	Email   string        `index:"unique" validate:"required,email"`
	Company string        `index:"-" validate:"min=5,max=40"`
	Address Address       `validate:"required,nested"`
}

func (c Customer) Marshal() map[string]interface{} {
	m := map[string]interface{}{"id": c.Id.Hex()}

	if len(c.Name) > 0 {
		m["name"] = c.Name
	}
	if len(c.Email) > 0 {
		m["email"] = c.Email
	}
	if len(c.Company) > 0 {
		m["company"] = c.Company
	}
	if a := c.Address.Marshal(); len(a) > 0 {
		m["address"] = a
	}

	return m
}

func (c Customer) MarshalJSON() ([]byte, error) {
	m := c.Marshal()

	return json.Marshal(m)
}
