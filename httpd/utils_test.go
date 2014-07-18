package httpd

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestError(t *testing.T) {
	Convey("Given a response and error nil", t, func() {
		res := httptest.NewRecorder()

		Convey("Error()", func() {
			Error(res, nil, http.StatusInternalServerError)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldContainSubstring, `{"error":"Internal Server Error"}`)
			})
		})
	})
	Convey("Given a response and an error", t, func() {
		res := httptest.NewRecorder()
		err := errors.New("Some error")

		Convey("Error()", func() {
			Error(res, err, http.StatusInternalServerError)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, http.StatusInternalServerError)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldContainSubstring, `{"error":"Some error"}`)
			})
		})
	})
}

func TestRespond(t *testing.T) {
	Convey("Given a response and value nil", t, func() {
		res := httptest.NewRecorder()

		Convey("Respond()", func() {
			Respond(res, nil, http.StatusNoContent)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, http.StatusNoContent)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldContainSubstring, `{"status":"No Content"}`)
			})
		})
	})
	Convey("Given a response and value map", t, func() {
		res := httptest.NewRecorder()

		Convey("Respond()", func() {
			Respond(res, map[string]string{"foo": "bar"}, http.StatusOK)

			Convey("Should set response status", func() {
				So(res.Code, ShouldEqual, http.StatusOK)
			})
			Convey("Should set response body", func() {
				So(res.Body.String(), ShouldContainSubstring, `{"foo":"bar"}`)
			})
		})
	})
}
