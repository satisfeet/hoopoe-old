package model

import "gopkg.in/mgo.v2/bson"

var (
	ProductName = "products"
)

type Product struct {
	Id          bson.ObjectId     `json:"id"     bson:"_id"`
	Name        string            `json:"name"`
	Pricing     Pricing           `json:"pricing"`
	Variations  ProductVariations `json:"variations"`
	Description string            `json:"description"`
}

type ProductVariations struct {
	Sizes  []string `json:"sizes"`
	Colors []string `json:"colors"`
}
