package store

import (
	"testing"

	"gopkg.in/mgo.v2/bson"

	. "github.com/smartystreets/goconvey/convey"
)

func TestQueryId(t *testing.T) {
	Convey("Given a valid string", t, func() {
		id := bson.NewObjectId()

		Convey("IdHex()", func() {
			query := Query{}
			query.Id(id.Hex())

			Convey("Should set _id", func() {
				So(query["_id"], ShouldEqual, id)
			})
		})
	})
	Convey("Given an invalid string", t, func() {
		id := "1234"

		Convey("IdHex()", func() {
			query := Query{}
			query.Id(id)

			Convey("Should not set _id", func() {
				So(query["_id"], ShouldBeNil)
			})
		})
	})
}

func TestQuerySearch(t *testing.T) {
	query := Query{}

	Convey("Given a string", t, func() {
		param := "Berlin"

		Convey("Search()", func() {
			query.Search(param, []string{"name", "email"})

			Convey("Should set or with regex", func() {
				r := bson.RegEx{"Berlin", "i"}

				So(query["$or"].([]bson.M)[0], ShouldResemble, bson.M{"name": r})
				So(query["$or"].([]bson.M)[1], ShouldResemble, bson.M{"email": r})
			})
		})
	})
}
