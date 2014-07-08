package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/satisfeet/hoopoe/store"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHttpdListen(t *testing.T) {
	Convey("Given a request to valid path", t, func() {
		req, res, httpd := request("GET", "/customers")

		Convey("ServeHTTP()", func() {
			httpd.ServeHTTP(res, req)

			Convey("Should respond: OK", func() {
				So(res.Code, ShouldEqual, 200)
			})
		})
	})
	Convey("Given a request with invalid path", t, func() {
		req, res, httpd := request("GET", "/")

		Convey("ServeHTTP()", func() {
			httpd.ServeHTTP(res, req)

			Convey("Should respond: Not Found", func() {
				So(res.Code, ShouldEqual, 404)
			})
		})
	})
}

func request(m string, p string) (*http.Request, *httptest.ResponseRecorder, *Httpd) {
	req, _ := http.NewRequest(m, p, nil)

	s := store.NewStore()
	s.Open("localhost/test")

	return req, httptest.NewRecorder(), NewHttpd(s)
}
