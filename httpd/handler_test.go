package httpd

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuth(t *testing.T) {
	Convey("Given a request without authorization header", t, func() {
		req, res := NewRequestResponse()

		Convey("Auth()", func() {
			Auth(HelloHandler()).ServeHTTP(res, req)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, http.StatusUnauthorized)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldEqual, http.StatusText(401))
			})
		})
	})
	Convey("Given a request with invalid authorization header", t, func() {
		req, res := NewRequestResponse()
		req.Header.Set("Authorization", "Basic")

		Convey("Auth()", func() {
			Auth(HelloHandler()).ServeHTTP(res, req)

			Convey("Should set response Unauthorized", func() {
				So(res.Code, ShouldEqual, http.StatusUnauthorized)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldEqual, http.StatusText(401))
			})
		})
	})
	Convey("Given a request with invalid authorization credentials", t, func() {
		req, res := NewRequestResponse()
		req.SetBasicAuth("foo", "bar")

		Convey("Auth()", func() {
			Auth(HelloHandler()).ServeHTTP(res, req)

			Convey("Should set response Unauthorized", func() {
				So(res.Code, ShouldEqual, http.StatusUnauthorized)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldEqual, http.StatusText(401))
			})
		})
	})
	Convey("Given a request with correct authorization credentials", t, func() {
		req, res := NewRequestResponse()
		req.SetBasicAuth(HttpUsername, HttpPassword)

		Convey("Auth()", func() {
			Auth(HelloHandler()).ServeHTTP(res, req)

			Convey("Should set response OK", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldEqual, "Hello World\n")
			})
		})
	})
}

func NewRequestResponse() (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", "/", nil)

	return req, httptest.NewRecorder()
}

func HelloHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello World\n"))
	})
}
