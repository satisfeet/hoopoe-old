package conf

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewConf(t *testing.T) {
	Convey("NewConf()", t, func() {
		conf := NewConf()

		Convey("Should return initialized Conf", func() {
			So(conf, ShouldHaveSameTypeAs, Conf{})
		})
	})
}

func TestConfParseFlags(t *testing.T) {
	Convey("Given no parameters", t, func() {
		conf := NewConf()

		Convey("ParseFlags()", func() {
			err := conf.ParseFlags()

			Convey("Should return error", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})
}
