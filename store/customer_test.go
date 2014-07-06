package store

import (
	"encoding/json"
	"fmt"
	"testing"

	"labix.org/v2/mgo/bson"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	Id = bson.NewObjectId()
)

const (
	CITY    = "Berlin"
	COMPANY = "satisfeet"
	EMAIL   = "bodo.kaiser@me.com"
	NAME    = "Bodo Kaiser"
	STREET  = "Geiserichstr. 3"
	ZIP     = 12105

	BASIC    = `{"id":"%s","name":"%s","email":"%s","address":{"city":"%s"}}`
	COMPLETE = `{"id":"%s","name":"%s","email":"%s","company":"%s","address":{"zip":%d,"city":"%s","street":"%s"}}`
)

func TestMarshalJSON(t *testing.T) {
	Convey("Given a basic model", t, func() {
		customer := Customer{
			Id:    Id,
			Name:  NAME,
			Email: EMAIL,
			Address: CustomerAddress{
				City: CITY,
			},
		}

		Convey("Which is marshaled as JSON", func() {
			json, err := json.Marshal(&customer)

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The result should be json", func() {
				So(string(json), ShouldEqual, fmt.Sprintf(BASIC, Id.Hex(), NAME, EMAIL, CITY))
			})
		})
	})
	Convey("Given a complete model", t, func() {
		customer := Customer{
			Id:      Id,
			Name:    NAME,
			Email:   EMAIL,
			Company: COMPANY,
			Address: CustomerAddress{
				Zip:    ZIP,
				City:   CITY,
				Street: STREET,
			},
		}

		Convey("Which is marshaled as JSON", func() {
			json, err := json.Marshal(&customer)

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The result should be json", func() {
				So(string(json), ShouldEqual, fmt.Sprintf(COMPLETE, Id.Hex(), NAME, EMAIL, COMPANY, ZIP, CITY, STREET))
			})
		})
	})
}

func TestUnmarshalJSON(t *testing.T) {
	j := fmt.Sprintf(COMPLETE, Id.Hex(), NAME, EMAIL, COMPANY, ZIP, CITY, STREET)

	Convey("Given a string", t, func() {
		customer := Customer{}

		Convey("Which is unmarshaled as JSON", func() {
			err := json.Unmarshal([]byte(j), &customer)

			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The customer should be complete", func() {
				So(customer.Id, ShouldEqual, Id)
				So(customer.Name, ShouldEqual, NAME)
				So(customer.Email, ShouldEqual, EMAIL)
				So(customer.Company, ShouldEqual, COMPANY)
				So(customer.Address.Zip, ShouldEqual, ZIP)
				So(customer.Address.City, ShouldEqual, CITY)
				So(customer.Address.Street, ShouldEqual, STREET)
			})
		})
	})
}
