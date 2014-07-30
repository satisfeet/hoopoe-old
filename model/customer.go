package model

import "gopkg.in/mgo.v2/bson"

type Customer struct {
	Id      bson.ObjectId `json:"id" bson:"_id"`
	Name    string        `json:"name,omitempty" validate:"required,min=5"`
	Email   string        `json:"email,omitempty" validate:"required,email"`
	Company string        `json:"company,omitempty" validate:"min=5,max=40"`
	Address Address       `json:"address,omitempty" validate:"required,nested"`
}
