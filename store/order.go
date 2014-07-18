package store

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	OrderIndices    = []string{}
	OrderCollection = "orders"
)

type Order struct {
	Id       bson.ObjectId `json:"id"     bson:"_id"`
	State    OrderState    `json:"state"`
	Items    []OrderItem   `json:"items"`
	Customer mgo.DBRef     `json:"customer"`
}

type OrderState struct {
	Created time.Time `json:"created"`
	Shipped time.Time `json:"shipped"`
	Cleared time.Time `json:"cleared"`
}

type OrderItem struct {
	Product   mgo.DBRef        `json:"product"`
	Price     float32          `json:"price"`
	Quantity  int              `json:"quantity"`
	Variation ProductVariation `json:"variation"`
}
