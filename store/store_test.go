package store

import (
	"testing"

	"labix.org/v2/mgo"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStoreOpen(t *testing.T) {
	store := NewStore()

	Convey("Given a valid url", t, func() {
		url := "mongodb://localhost/test"

		Convey("Open()", func() {
			err := store.Open(url)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
	// this test case slows takes up to 10s because
	// mgo waits up its timeout for searching servers
	// because of this we will skip this for now
	SkipConvey("Given an invalid url", t, func() {
		url := "http://localhost:2000"

		Convey("Open()", func() {
			err := store.Open(url)

			Convey("Should return error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}

func TestStoreMongo(t *testing.T) {
	store := NewStore()

	Convey("Given an opened store", t, func() {
		store.Open("localhost/test")

		Convey("Mongo()", func() {
			mongo := store.Mongo()

			Convey("Should return mgo.Database", func() {
				So(mongo, ShouldHaveSameTypeAs, &mgo.Database{})
			})
		})
	})
	Convey("Given an unopened store", t, func() {
		store.Close()

		Convey("Mongo()", func() {
			Convey("Should panic", func() {
				So(func() { store.Mongo() }, ShouldPanic)
			})
		})
	})
}

func TestStoreClose(t *testing.T) {
	store := NewStore()

	Convey("Given an opened store", t, func() {
		store.Open("localhost/test")

		Convey("Close()", func() {
			So(store.Close, ShouldNotPanic)
		})
	})
}
