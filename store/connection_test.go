package store

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestOpen(t *testing.T) {
	Convey("Given a valid url", t, func() {
		url := "mongodb://localhost/test"

		Convey("Open()", func() {
			err := Open(url)

			Convey("Should return no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("Should set session", func() {
				So(mongo, ShouldNotBeNil)
			})
		})
	})
}

func TestClose(t *testing.T) {
	Convey("Given an opened store", t, func() {
		Open("localhost/test")

		Convey("Close()", func() {
			So(Close, ShouldNotPanic)

			Convey("Should set session", func() {
				So(mongo, ShouldBeNil)
			})
		})
	})
}
