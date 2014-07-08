package main

import (
	"net/http"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {
	Convey("Given arguments", t, func() {
		os.Args = append(os.Args,
			"--addr", "127.0.0.1:3001",
			"--mongo", "localhost/test",
		)

		Convey("main()", func() {
			// do not block tests
			go main()

			// "sync" with main routine
			//
			// this is necessary as requests will come in
			// before the http server listens on a tcp socket
			time.Sleep(time.Second)

			Convey("Should handle HTTP requests", func() {
				res, err := http.Get("http://127.0.0.1:3001/customers")

				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 200)
			})
		})
	})
}
