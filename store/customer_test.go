package store

import (
	"encoding/json"
	"fmt"
	"testing"

	"labix.org/v2/mgo/bson"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	id      = bson.NewObjectId()
	name    = "Bodo Kaiser"
	email   = "bodo.kaiser@me.com"
	company = "satisfeet"
	street  = "Geiserichstr. 3"
	city    = "Berlin"
	zip     = 12105
)

func TestCustomerMarshalJSON(t *testing.T) {
	Convey("Given basic struct data", t, func() {
		customer := Customer{
			Id:    id,
			Name:  name,
			Email: email,
			Address: CustomerAddress{
				City: city,
			},
		}

		Convey("MarshalJSON()", func() {
			json, err := json.Marshal(&customer)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should return struct as json", func() {
				So(string(json), ShouldEqual, formatBasic())
			})
		})
	})
	Convey("Given advanced struct data", t, func() {
		customer := Customer{
			Id:      id,
			Name:    name,
			Email:   email,
			Company: company,
			Address: CustomerAddress{
				Zip:    zip,
				City:   city,
				Street: street,
			},
		}

		Convey("MarshalJSON()", func() {
			json, err := json.Marshal(&customer)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should return struct as json", func() {
				So(string(json), ShouldEqual, formatAdvanced())
			})
		})
	})
}

func TestCustomerUnmarshalJSON(t *testing.T) {
	j := formatAdvanced()

	Convey("Given a json string", t, func() {
		customer := Customer{}

		Convey("UnmarshalJSON()", func() {
			err := json.Unmarshal([]byte(j), &customer)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should set customer id", func() {
				So(customer.Id, ShouldEqual, id)
			})
			Convey("Should set customer name", func() {
				So(customer.Name, ShouldEqual, name)
			})
			Convey("Should set customer email", func() {
				So(customer.Email, ShouldEqual, email)
			})
			Convey("Should set customer company", func() {
				So(customer.Company, ShouldEqual, company)
			})
			Convey("Should set customer zip address", func() {
				So(customer.Address.Zip, ShouldEqual, zip)
			})
			Convey("Should set customer city address", func() {
				So(customer.Address.City, ShouldEqual, city)
			})
			Convey("Should set customer street address", func() {
				So(customer.Address.Street, ShouldEqual, street)
			})
		})
	})
}

func formatBasic() string {
	s := `{"id":"%s","name":"%s","email":"%s","address":{"city":"%s"}}`

	return fmt.Sprintf(s, id.Hex(), name, email, city)
}

func formatAdvanced() string {
	s := `{"id":"%s","name":"%s","email":"%s","company":"%s","address":{"zip":%d,"city":"%s","street":"%s"}}`

	return fmt.Sprintf(s, id.Hex(), name, email, company, zip, city, street)
}
