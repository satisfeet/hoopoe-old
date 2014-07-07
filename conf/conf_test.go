package conf

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {
	Convey("New()", t, func() {
		conf := New()

		Convey("Should return initialized Conf", func() {
			So(conf, ShouldHaveSameTypeAs, Conf{})
		})
	})
}

func TestConfParseFlags(t *testing.T) {
	Convey("Given no parameters", t, func() {
		conf := New()

		Convey("ParseFlags()", func() {
			err := conf.ParseFlags()

			Convey("Should return error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
