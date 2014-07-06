package store

import (
	"encoding/json"
	"fmt"
	"testing"

	"labix.org/v2/mgo/bson"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	id      = bson.NewObjectid()
	name    = "Bodo Kaiser"
	email   = "bodo.kaiser@me.com"
	company = "satisfeet"
	street  = "Geiserichstr. 3"
	city    = "Berlin"
	zip     = 12105
)

const (
	BASIC    = `{"id":"%s","name":"%s","email":"%s","address":{"city":"%s"}}`
	COMPLETE = `{"id":"%s","name":"%s","email":"%s","company":"%s","address":{"zip":%d,"city":"%s","street":"%s"}}`
)

func TestMarshalJSON(t *testing.T) {
	Convey("Given a basic model", t, func() {
		customer := Customer{
			id:    id,
			Name:  name,
			Email: email,
			Address: CustomerAddress{
				City: city,
			},
		}

		Convey("Which is marshaled as JSON", func() {
			json, err := json.Marshal(&customer)

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The result should be json", func() {
				So(string(json), ShouldEqual, fmt.Sprintf(BASIC,
					id.Hex(), name, email, city))
			})
		})
	})
	Convey("Given a complete model", t, func() {
		customer := Customer{
			id:      id,
			Name:    name,
			Email:   email,
			Company: company,
			Address: CustomerAddress{
				Zip:    zip,
				City:   city,
				Street: street,
			},
		}

		Convey("Which is marshaled as JSON", func() {
			json, err := json.Marshal(&customer)

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The result should be json", func() {
				So(string(json), ShouldEqual, fmt.Sprintf(COMPLETE,
					id.Hex(), name, email, company, zip, city, street))
			})
		})
	})
}

func TestUnmarshalJSON(t *testing.T) {
	j := fmt.Sprintf(COMPLETE, id.Hex(),
		name, email, company, zip, city, street)

	Convey("Given a string", t, func() {
		customer := Customer{}

		Convey("Which is unmarshaled as JSON", func() {
			err := json.Unmarshal([]byte(j), &customer)

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The customer should be complete", func() {
				So(customer.id, ShouldEqual, id)
				So(customer.Name, ShouldEqual, name)
				So(customer.Email, ShouldEqual, email)
				So(customer.Company, ShouldEqual, company)
				So(customer.Address.Zip, ShouldEqual, zip)
				So(customer.Address.City, ShouldEqual, city)
				So(customer.Address.Street, ShouldEqual, street)
			})
		})
	})
}
