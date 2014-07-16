package store

import "labix.org/v2/mgo/bson"

var (
	ProductDatabase   = ""
	ProductCollection = "products"

	ProductUnique = []string{
		"name",
	}
)

type Product struct {
	Id          bson.ObjectId     `json:"id"     bson:"_id"`
	Name        string            `json:"name"`
	Pricing     ProductPricing    `json:"pricing"`
	Variations  ProductVariations `json:"variations"`
	Description string            `json:"description"`
}

type ProductPricing struct {
	Retail float32 `json:"retail"`
}

type ProductVariations struct {
	Sizes  []string `json:"sizes"`
	Colors []string `json:"colors"`
}
