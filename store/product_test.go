package store

import (
	"bytes"
	"encoding/json"
	"testing"

	"gopkg.in/mgo.v2/bson"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	// Product as go struct.
	ProductStruct = &Product{
		Id:   bson.ObjectIdHex("1234567890abcd1234567890"),
		Name: "Summer Socks",
		Pricing: ProductPricing{
			Retail: 2.99,
		},
		Variations: ProductVariations{
			Sizes:  []string{"42-44", "46-48"},
			Colors: []string{"blue", "black"},
		},
		Description: "These are nice summer socks.",
	}

	// Product as json string.
	ProductJSON = []byte(`{
		"id": "1234567890abcd1234567890",
		"name": "Summer Socks",
		"pricing": {
			"retail": 2.99
		},
		"variations": {
			"sizes": ["42-44", "46-48"],
			"colors": ["blue", "black"]
		},
		"description": "These are nice summer socks."
	}`)
)

func TestProductJSON(t *testing.T) {
	Convey("Given product as json", t, func() {
		data := ProductJSON

		Convey("json.Unmarshal()", func() {
			product := &Product{}
			err := json.Unmarshal(data, product)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should set product fields", func() {
				So(product, ShouldResemble, ProductStruct)
			})
		})
	})
	Convey("Given product as struct", t, func() {
		product := *ProductStruct

		Convey("json.Marshal()", func() {
			data, err := json.Marshal(product)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should return product as json", func() {
				buf := new(bytes.Buffer)
				json.Compact(buf, ProductJSON)

				So(data, ShouldResemble, buf.Bytes())
			})
		})
	})
}
