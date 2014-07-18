package store

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
)

var (
	// Our test url.
	url = "localhost/test"
	// Our test name.
	name = "testers"
)

type (
	// Our test model.
	model struct {
		Id   bson.ObjectId `bson:"_id"`
		Text string
	}
)

func TestStore(t *testing.T) {
	Open(url)
	defer Close()

	Convey("Given a value and a store", t, func() {
		value := model{bson.NewObjectId(), "I am a tester!"}

		Convey("Insert()", func() {
			err := NewStore(name).Insert(&value)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should save value", func() {
				count, _ := mongo.DB(Database).C(name).Find(nil).Count()

				So(count, ShouldEqual, 1)
			})
			Reset(func() {
				mongo.DB(Database).C(name).DropCollection()
			})
		})
	})
	Convey("Given a stored value", t, func() {
		value := model{bson.NewObjectId(), "Hey ho!"}
		mongo.DB(Database).C(name).Insert(&value)
		value.Text += "?"

		Convey("Update()", func() {
			err := NewStore(name).Update(Query{"_id": value.Id}, &value)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should save value", func() {
				mongo.DB(Database).C(name).FindId(value.Id).One(&value)

				So(value.Text, ShouldEqual, "Hey ho!?")
			})
			Reset(func() {
				mongo.DB(Database).C(name).DropCollection()
			})
		})
	})
	Convey("Given a stored value", t, func() {
		value := model{bson.NewObjectId(), "Foobar"}
		mongo.DB(Database).C(name).Insert(&value)

		Convey("Remove()", func() {
			err := NewStore(name).Remove(Query{"_id": value.Id})

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should remove value", func() {
				count, _ := mongo.DB(Database).C(name).Find(nil).Count()

				So(count, ShouldEqual, 0)
			})
			Reset(func() {
				mongo.DB(Database).C(name).DropCollection()
			})
		})
	})
}
